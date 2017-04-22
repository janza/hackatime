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
)

var (
	bucketName = []byte("LogBucket")
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

func handleReadRequest(r io.ReadCloser) (readRequest, error) {
	req := readRequest{}
	decoder := json.NewDecoder(r)
	err := decoder.Decode(&req)
	if err != nil && err != io.EOF {
		return req, err
	}
	now := time.Now()
	if req.Min.IsZero() {
		if req.Max.IsZero() {
			req.Max = now
		}
		req.Min = req.Max.Add(-time.Hour)
	}
	if req.Max.IsZero() {
		req.Max = req.Min.Add(time.Hour)
	}

	return req, nil
}

type app struct {
	db *bolt.DB
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

func (a *app) reader(w http.ResponseWriter, r *http.Request) {

	bounds, err := handleReadRequest(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusServiceUnavailable)
		fmt.Fprintf(w, "Error parsing request: %s", err)
		return
	}

	err = a.db.View(func(tx *bolt.Tx) error {
		c := tx.Bucket(bucketName).Cursor()
		min := bounds.dbFormatMin()
		max := bounds.dbFormatMax()

		log.Printf("Loading from %s to %s", min, max)

		for k, v := c.Seek(min); k != nil && bytes.Compare(k, max) <= 0; k, v = c.Next() {
			var entry logInput
			err := json.Unmarshal(v, &entry)
			if err != nil {
				log.Printf("error decoding stored value: %s", err)
				continue
			}
			fmt.Fprintf(w, "%s\n", string(v))
		}

		return nil
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

func newApp() (*app, error) {
	a := &app{}
	err := a.initDb()
	return a, err
}

func main() {

	a, err := newApp()
	if err != nil {
		log.Fatal(err)
	}
	defer a.cleanup()
	http.HandleFunc("/log", a.writer)
	http.HandleFunc("/read", a.reader)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Fatal(http.ListenAndServe(":"+port, nil))
}
