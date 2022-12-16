package challenges

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay16ParseLine(t *testing.T) {
	d := Day16{}
	n, v := d.parseLine("Valve DD has flow rate=20; tunnels lead to valves CC, AA, EE")
	assert.Equal(t, "DD", n)
	assert.Equal(t, Valve{FlowRate: 20, Valves: []string{"CC", "AA", "EE"}}, v)

	n, v = d.parseLine("Valve JJ has flow rate=21; tunnel leads to valve II")
	assert.Equal(t, "JJ", n)
	assert.Equal(t, Valve{FlowRate: 21, Valves: []string{"II"}}, v)
}

func TestDay16Load(t *testing.T) {
	d := Day16{}
	assert.NoError(t, d.Load("testfiles/day16.txt"))
	assert.Len(t, d.vs, 10)
}

func TestDay16CreateAndOptimize(t *testing.T) {
	d := Day16{}
	assert.NoError(t, d.Load("testfiles/day16.txt"))
	assert.Len(t, d.vs, 10)
	vs2 := d.optimizeSystem(d.vs)
	assert.NotEqual(t, d.vs, vs2)
	assert.Equal(t, d.vs["CC"], vs2["CC"])
	assert.NotEqual(t, d.vs["BB"], vs2["BB"])
	assert.Len(t, d.vs["BB"].Valves, 2)
	assert.Len(t, vs2["BB"].Valves, 1)

}

func TestDay16p1(t *testing.T) {
	d := Day16{}
	assert.NoError(t, d.Load("testfiles/day16.txt"))
	assert.Equal(t, "1651", d.Part1())
}
