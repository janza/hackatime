package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/boltdb/bolt"
	"github.com/gorilla/websocket"
)

var (
	bucketName = []byte("LogBucket")
	upgrader   = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
)

type logInput struct {
	Entity   string    `json:"entity"`
	Time     time.Time `json:"time"`
	Project  string    `json:"project"`
	Branch   string    `json:"branch"`
	Language string    `json:"language"`
}

type readRequest struct {
	Min time.Time
	Max time.Time
}

func (r *readRequest) dbFormatMin() []byte {
	return []byte(r.Min.UTC().Format(time.RFC3339))
}

func (r *readRequest) dbFormatMax() []byte {
	return []byte(r.Max.UTC().Format(time.RFC3339))
}

func (r *readRequest) setDefaults() {
	now := time.Now()
	if r.Min.IsZero() {
		if r.Max.IsZero() {
			r.Max = now
		}
		r.Min = r.Max.Add(-time.Hour)
	}
	if r.Max.IsZero() {
		r.Max = r.Min.Add(time.Hour)
	}
}

func handleReadRequest(r io.ReadCloser) (readRequest, error) {
	req := readRequest{}
	decoder := json.NewDecoder(r)
	err := decoder.Decode(&req)
	if err != nil && err != io.EOF {
		return req, err
	}
	req.setDefaults()
	return req, nil
}

type app struct {
	db   *bolt.DB
	send chan []byte
	hub  *hub
}

func (a *app) writer(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Use POST")
		return
	}
	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()
	var t logInput
	err := decoder.Decode(&t)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Invalid JSON %s", err)
		return
	}

	s, err := json.Marshal(t)
	if err != nil {
		w.WriteHeader(http.StatusServiceUnavailable)
		fmt.Fprintf(w, "Error converting to JSON: %s", err)
		return
	}

	a.hub.broadcast <- s

	err = a.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(bucketName)
		if err != nil {
			return err
		}
		log.Printf("Writing: %s", s)
		return b.Put([]byte(t.Time.UTC().Format(time.RFC3339)), []byte(s))
	})
	if err != nil {
		w.WriteHeader(http.StatusServiceUnavailable)
		fmt.Fprintf(w, "Error writing to database: %s", err)
		return
	}
	fmt.Fprintf(w, `{"ok": %s}`, s)
}

func (a *app) getMessages(r readRequest, messageHandler func([]byte)) error {
	return a.db.View(func(tx *bolt.Tx) error {
		c := tx.Bucket(bucketName).Cursor()
		min := r.dbFormatMin()
		max := r.dbFormatMax()

		log.Printf("Loading from %s to %s", min, max)

		for k, v := c.Seek(min); k != nil && bytes.Compare(k, max) <= 0; k, v = c.Next() {
			messageHandler(v)
		}

		return nil
	})
}

func (a *app) reader(w http.ResponseWriter, r *http.Request) {

	bounds, err := handleReadRequest(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusServiceUnavailable)
		fmt.Fprintf(w, "Error parsing request: %s", err)
		return
	}

	err = a.getMessages(bounds, func(m []byte) {
		fmt.Fprintf(w, "%s\n", string(m))
	})
	if err != nil {
		w.WriteHeader(http.StatusServiceUnavailable)
		fmt.Fprintf(w, "Error parsing request: %s", err)
		return
	}
}

func (a *app) initDb() error {
	db, err := bolt.Open("my.db", 0600, nil)
	if err != nil {
		return err
	}

	err = db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists(bucketName)
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
	a.db = db
	return nil
}

func (a *app) cleanup() error {
	return a.db.Close()
}

func (a *app) serveWs(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	c := &client{
		hub:  a.hub,
		conn: conn,
		send: make(chan []byte),
	}
	a.hub.register <- c
	ticker := time.NewTicker(50 * time.Second)
	defer func() {
		ticker.Stop()
		a.hub.unregister <- c
		conn.Close()
	}()

	go func() {
		req := readRequest{}
		n := time.Now()
		req.Max = n
		req.Min = n.Add(time.Hour * -1)
		a.getMessages(req, func(m []byte) {
			a.hub.broadcast <- m
		})
	}()

	for {
		select {
		case message, ok := <-c.send:
			conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
			if !ok {
				// The hub closed the channel.
				conn.WriteMessage(websocket.CloseMessage, []byte{})
				log.Println("hub closed channel")
				return
			}

			w, err := conn.NextWriter(websocket.TextMessage)
			if err != nil {
				log.Printf("err on next writer: %s", err)
				return
			}
			log.Println("writing message")
			w.Write(message)
			log.Println("closing writer")
			if err := w.Close(); err != nil {
				log.Printf("err on close: %s", err)
				return
			}
		case <-ticker.C:
			conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
			if err := conn.WriteMessage(websocket.PingMessage, []byte{}); err != nil {
				return
			}
		}
	}
}

func newApp() (*app, error) {
	a := &app{}
	a.send = make(chan []byte)
	a.hub = &hub{
		broadcast:  make(chan []byte),
		register:   make(chan *client),
		unregister: make(chan *client),
		clients:    make(map[*client]bool),
	}
	go a.hub.run()
	err := a.initDb()
	return a, err
}

type client struct {
	hub  *hub
	conn *websocket.Conn
	send chan []byte
}

type hub struct {
	clients    map[*client]bool
	broadcast  chan []byte
	register   chan *client
	unregister chan *client
}

func newHub() *hub {
	return &hub{
		broadcast:  make(chan []byte),
		register:   make(chan *client),
		unregister: make(chan *client),
		clients:    make(map[*client]bool),
	}
}

func (h *hub) run() {

	for {
		select {
		case client := <-h.register:
			h.clients[client] = true
		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.send)
			}
		case message := <-h.broadcast:
			for client := range h.clients {
				client.send <- message
			}
		}
	}
}

//go:generate go-bindata -pkg main -o bindata.go static

func staticHandler(rw http.ResponseWriter, req *http.Request) {
	path := req.URL.Path[1:]
	if path == "" {
		path = "index.html"
	}
	if bs, err := Asset(path); err != nil {
		log.Printf("path not found: %s", path)
		rw.WriteHeader(http.StatusNotFound)
	} else {
		var reader = bytes.NewBuffer(bs)
		io.Copy(rw, reader)
	}
}

func main() {

	a, err := newApp()
	if err != nil {
		log.Fatal(err)
	}
	defer a.cleanup()
	http.HandleFunc("/log", a.writer)
	http.HandleFunc("/read", a.reader)
	http.HandleFunc("/ws", a.serveWs)
	http.Handle("/static/", http.HandlerFunc(staticHandler))

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Fatal(http.ListenAndServe(":"+port, nil))
}
