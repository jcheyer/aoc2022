package challenges

import (
	"fmt"
	"strings"

	"github.com/RyanCarrier/dijkstra"
	"github.com/jcheyer/aoc2021/lib"
)

type Day16 struct {
	rawInput  []string
	vs        ValveSystem
	graph     *dijkstra.Graph
	distances map[string]int
}

type Valve struct {
	FlowRate int
	Valves   []string
	Open     bool
}
type ValveSystem map[string]Valve

func (d *Day16) Load(file string) error {
	if len(d.rawInput) > 0 {
		return nil
	}
	var err error
	d.rawInput, err = lib.ReadInputLinesRaw(file)
	if err != nil {
		return err
	}
	d.vs = d.createValveSystem()
	d.graph = d.createGraph()
	d.distances = d.precalcDistances()

	//d.vs = d.optimizeSystem(d.vs)

	return nil
}

func (d Day16) createValveSystem() ValveSystem {
	vs := make(ValveSystem)

	for _, line := range d.rawInput {
		name, v := d.parseLine(line)
		vs[name] = v
	}

	return vs
}

func (d Day16) createGraph() *dijkstra.Graph {
	graph := dijkstra.NewGraph()

	for name, _ := range d.vs {
		graph.AddMappedVertex(name)
	}

	for k, v := range d.vs {
		for _, valve := range v.Valves {
			err := graph.AddMappedArc(k, valve, 1)
			if err != nil {
				panic(err.Error())
			}
		}
	}

	return graph
}

func (d Day16) precalcDistances() map[string]int {
	distances := map[string]int{}
	for name1 := range d.vs {
		for name2 := range d.vs {
			id1, _ := d.graph.GetMapping(name1)
			id2, _ := d.graph.GetMapping(name2)
			path, _ := d.graph.Shortest(id1, id2)
			distances[name1+name2] = int(path.Distance)
		}
	}

	return distances
}

func (d *Day16) optimizeSystem(vs ValveSystem) ValveSystem {

	vsNew := make(ValveSystem)

	for name, valve := range vs {
		useFullTunnels := []string{}
		for _, v := range valve.Valves {
			if d.vs[v].FlowRate != 0 {
				useFullTunnels = append(useFullTunnels, v)
			}
		}
		valve.Valves = useFullTunnels
		vsNew[name] = valve
	}
	return vsNew
}

func (d Day16) parseLine(line string) (string, Valve) {
	l := strings.ReplaceAll(line, ", ", ",")
	l = strings.ReplaceAll(l, "tunnel leads to valve ", "tunnels lead to valves ")
	v := Valve{}
	name := ""
	tmp := ""
	fmt.Sscanf(l, "Valve %s has flow rate=%d; tunnels lead to valves %s", &name, &v.FlowRate, &tmp)
	v.Valves = strings.Split(tmp, ",")
	return name, v
}

func (d Day16) solveP1(step, pressure, flow int, currentValve string, tunnels []string, limit int, depth int, debug bool) int {

	maxPressure := pressure + (limit-step)*flow

	if debug {
		fmt.Printf("%s Step: %d Pressure: %d Flow: %d CurrentValve: %s\n", strings.Repeat("  ", depth), step, pressure, flow, currentValve)
		fmt.Printf("%s maxPressure: %d\n", strings.Repeat("  ", depth), maxPressure)
		fmt.Printf("%s ToDo: %+v\n", strings.Repeat("  ", depth), tunnels)
	}
	for _, tunnel := range tunnels {
		/*currentValveID, err := d.graph.GetMapping(currentValve)
		if err != nil {
			panic("CurrentValve...." + err.Error())
		}
		nextValveID, err := d.graph.GetMapping(tunnel)
		if err != nil {
			panic("NextValve...." + err.Error())
		}

		walk, err := d.graph.Shortest(currentValveID, nextValveID)
		if err != nil {
			panic("Shortest... " + err.Error() + "\n" + tunnel + " " + currentValve)
		}

		walkAndOpen := int(walk.Distance + 1)*/
		walkAndOpen := d.distances[currentValve+tunnel] + 1
		if step+walkAndOpen < limit {
			nextStep := step + walkAndOpen
			nextPressure := pressure + walkAndOpen*flow
			nextFlow := flow + d.vs[tunnel].FlowRate
			x := d.solveP1(nextStep, nextPressure, nextFlow, tunnel, d.cut(tunnels, tunnel), limit, depth+1, debug)
			if x > maxPressure {
				maxPressure = x
			}
		}
	}
	if debug {
		fmt.Printf("%s ------ returning %d ------ \n", strings.Repeat("  ", depth), maxPressure)
	}
	return maxPressure
}

func (d Day16) cutDefunc(s []string, i int) []string {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

func (d Day16) cut(s []string, delete string) []string {
	ret := []string{}
	for _, i := range s {
		if i != delete {
			ret = append(ret, i)
		}
	}
	return ret
}

func (d *Day16) Part1() string {
	res := 0

	toOpen := []string{}
	for name, valve := range d.vs {
		if valve.FlowRate > 0 {
			toOpen = append(toOpen, name)
		}
	}

	res = d.solveP1(0, 0, 0, "AA", toOpen, 30, 0, false)

	return fmt.Sprintf("%d", res)
}

func (d *Day16) Part2() string {
	res := 0

	return fmt.Sprintf("%d", res)
}
