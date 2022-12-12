package mountainmap

import (
	"github.com/RyanCarrier/dijkstra"
)

type coord struct {
	x, y int
}

type Mountainmap struct {
	D      *dijkstra.Graph
	data   map[coord]int
	p2     []coord
	start  coord
	target coord
	dimX   int
	dimY   int
}

func New(data []string) (*Mountainmap, error) {
	mm := &Mountainmap{
		D:    dijkstra.NewGraph(),
		p2:   make([]coord, 0),
		data: make(map[coord]int),
		dimX: len(data[0]),
		dimY: len(data),
	}

	for y, line := range data {
		for x, r := range line {

			if r == 'S' {
				mm.start = coord{x, y}
				r = 'a'
			}
			if r == 'E' {
				mm.target = coord{x, y}
				r = 'z'
			}
			if r == 'a' {
				mm.p2 = append(mm.p2, coord{x, y})
			}
			mm.data[coord{x, y}] = int(r)
			vertexID := mm.Coord2vertex(x, y)
			mm.D.AddVertex(vertexID)
		}
	}

	for y := 0; y < mm.dimY; y++ {
		for x := 0; x < mm.dimX; x++ {
			vertexID1 := mm.Coord2vertex(x, y)
			neighbors := mm.reachableAround(x, y)
			for id, _ := range neighbors {
				vertexID2 := mm.Coord2vertex(id.x, id.y)
				err := mm.D.AddArc(vertexID1, vertexID2, 1)
				if err != nil {
					return mm, err
				}
			}
		}
	}

	return mm, nil
}
func (mm *Mountainmap) Part1() int64 {

	v1 := mm.Coord2vertex(mm.start.x, mm.start.y)
	v2 := mm.Coord2vertex(mm.target.x, mm.target.y)

	path, err := mm.D.Shortest(v1, v2)
	if err != nil {
		panic(err.Error())
	}

	return path.Distance
}

func (mm *Mountainmap) Part2() int64 {

	shortest := int64(10000000)

	for _, c := range mm.p2 {
		v1 := mm.Coord2vertex(c.x, c.y)
		v2 := mm.Coord2vertex(mm.target.x, mm.target.y)

		path, err := mm.D.Shortest(v1, v2)
		if err != nil {
			continue

		}
		if path.Distance < shortest {
			shortest = path.Distance
		}

	}

	return shortest
}

func (mm *Mountainmap) Coord2vertex(x, y int) int {
	return mm.dimX*(y) + x
}

func (mm *Mountainmap) reachableAround(x, y int) map[coord]int {
	height := mm.data[coord{x, y}]
	neighbors := make(map[coord]int)

	val, ok := mm.data[coord{x - 1, y}]
	if ok {
		if reachable(val, height) {
			neighbors[coord{x - 1, y}] = val
		}
	}

	val, ok = mm.data[coord{x + 1, y}]
	if ok {
		if reachable(val, height) {
			neighbors[coord{x + 1, y}] = val
		}
	}

	val, ok = mm.data[coord{x, y - 1}]
	if ok {
		if reachable(val, height) {
			neighbors[coord{x, y - 1}] = val
		}
	}

	val, ok = mm.data[coord{x, y + 1}]
	if ok {
		if reachable(val, height) {
			neighbors[coord{x, y + 1}] = val
		}
	}

	return neighbors
}

func reachable(a, b int) bool {
	if a <= b {
		return true
	}
	if a-1 == b {
		return true
	}

	return false
}
