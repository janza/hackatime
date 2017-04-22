var d3 = window.d3

const socketProto = window.location.protocol === 'https:' ? 'wss' : 'ws'
const socket = new window.WebSocket(`${socketProto}://${window.location.host}/ws`)

var lastTime = {}
var totalTime = {}
var firstTime = {}

function calcTime (project, item) {
  const d = +new Date(item.time)
  if (!lastTime[project]) {
    lastTime[project] = d
    firstTime[project] = d
  }
  const diff = d - lastTime[project]

  if (diff < 1000 * 60 * 5) {
    totalTime[project] = (totalTime[project] || 0) + diff
  }

  lastTime[project] = d
  return totalTime[project]
}

var updateState = debounce((items) => {
  lastTime = {}
  totalTime = {}

  const o = items.reduce((s, i) => {
    const project = i.project || 'other'
    s[project] = calcTime(project, i)
    return s
  }, {})
  render(
  Object.keys(o).sort((a, b) => a.localeCompare(b)).map(i => ({
    value: o[i],
    name: i
  })).filter(i => !!i.value)
  )
}, 100)

// Listen for messages
const items = []
socket.addEventListener('message', function (event) {
  const item = JSON.parse(event.data)
  items.push(item)
  updateState(items)
})

// d3

var width = 500
var height = 500
var radius = Math.min(width, height) / 2 - 10

var color = d3.scaleOrdinal(d3.schemeCategory20c)

var svg = d3.select('svg')

var g = svg
  .append('g')
  .attr('transform', 'translate(' + width / 2 + ',' + height / 2 + ')')

var pie = d3
  .pie()
  .padAngle(0.01)
  .value(d => d.value)
  .sort((a, b) => a.name.localeCompare(b.name))

var path = d3
  .arc()
  .outerRadius(radius)
  .innerRadius(Math.min(width, height) / 10)
  .cornerRadius(7)

var label = d3
  .arc()
  .outerRadius(radius - 40)
  .innerRadius(radius - 40)

function r (data) {
  if (data.length < 2) return

  var arc = g.selectAll('.arc')
    .data(pie(data), d => d.data.name)

  var enterArc = arc
    .enter()
    .append('g')
    .attr('class', 'arc')

  enterArc.append('path')
    .attr('d', path)
    .attr('fill', d => color(d.data.name))

  enterArc.append('text')
    .attr('transform', (d) => 'translate(' + label.centroid(d) + ')')
    .attr('dy', '0.35em')
    .text(d => `${d.data.name} (${(d.data.value / 1000 / 60 / 60).toFixed(2)})`)

  arc.exit().remove()

  arc
    .select('path')
    .transition()
    .duration(500)
    .attr('d', path)
    .attr('fill', d => color(d.data.name))

  arc
    .select('text')
    .attr('transform', (d) => 'translate(' + label.centroid(d) + ')')
    .attr('dy', '0.35em')
    .text(d => d.data.name)
}

function debounce (fn, time) {
  var timeout
  return function (...args) {
    clearTimeout(timeout)
    timeout = setTimeout(() => fn(...args), time)
  }
}

var render = debounce(r, 100)
