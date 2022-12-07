package directory

import (
	"testing"

	"github.com/jcheyer/aoc2021/lib"
	"github.com/stretchr/testify/assert"
)

func TestDirectory(t *testing.T) {
	root := New("/")
	assert.Equal(t, 0, root.Size())
	root.AddFile(NewFile("bla", 100))
	assert.Equal(t, 100, root.Size())
	bla := New("bla")
	bla.AddFile(NewFile("bla2", 200))
	root.AddDir(bla)
	assert.Equal(t, 200, bla.Size())
	assert.Equal(t, 300, root.Size())
}

func TestParse(t *testing.T) {
	data, err := lib.ReadInputLines("testfiles/day07.txt")
	assert.NoError(t, err)
	dir := ParseTerminal(data)
	assert.NotNil(t, dir)
	assert.Equal(t, 48381165, dir.Size())
	assert.Equal(t, 584, dir.Subdirs["a"].Subdirs["e"].Size())
	assert.Equal(t, 94853, dir.Subdirs["a"].Size())
	assert.Equal(t, 24933642, dir.Subdirs["d"].Size())
	assert.Equal(t, 95437, dir.SumSizeOfAllDirsSmallerThan(100000))
}
