var d3 = window.d3

const socketProto = window.location.protocol === 'https:' ? 'wss' : 'ws'
const socket = new window.WebSocket(`${socketProto}://${window.location.host}/ws`)

// Listen for messages
const items = []
socket.addEventListener('message', function (event) {
  const item = JSON.parse(event.data)
  items.push(item)
  const o = items.reduce((s, i) => {
    s[i.project] = (s[i.project] || 0) + 1
    return s
  }, {})
  render(
    Object.keys(o).sort((a, b) => a.localeCompare(b)).map(i => ({
      value: o[i],
      name: i
    }))
  )
})

// d3

var width = window.innerWidth
var height = window.innerHeight
var radius = Math.min(width, height) / 2 - 10

var color = d3.scaleOrdinal(d3.schemeCategory20c)

var svg = d3.select('svg')

var g = svg
  .append('g')
  .attr('transform', 'translate(' + width / 2 + ',' + height / 2 + ')')

var pie = d3
  .pie()
  .value(d => d.value)
  .sort((a, b) => a.name.localeCompare(b.name))

var path = d3
  .arc()
  .outerRadius(radius)
  .innerRadius(Math.min(width, height) / 10)

function r (data) {
  if (data.length < 2) return

  var arc = g.selectAll('.arc').data(pie(data))

  arc
    .enter()
    .append('g')
    .attr('class', 'arc')
    .append('path')
    .attr('d', path)
    .attr('fill', d => color(d.data.name))

  arc.exit().remove()

  arc
    .select('path')
    .transition()
    .duration(500)
    .attr('d', path)
    .attr('fill', d => color(d.data.name))
}

var timeout
function render (data) {
  clearTimeout(timeout)
  timeout = setTimeout(() => r(data), 500)
}
