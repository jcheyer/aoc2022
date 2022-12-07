package challenges

import (
	"fmt"
	"sort"

	"github.com/jcheyer/aoc2021/challenges/internal/directory"
	"github.com/jcheyer/aoc2021/lib"
)

type Day07 struct {
	rawInput []string
	dir      *directory.Directory
}

func (d *Day07) Load(file string) error {

	if len(d.rawInput) > 0 {
		return nil
	}
	var err error
	d.rawInput, err = lib.ReadInputLines(file)
	if err != nil {
		return err
	}

	d.dir = directory.ParseTerminal(d.rawInput)

	return nil
}

func (d *Day07) Part1() string {

	return fmt.Sprintf("%d", d.dir.SumSizeOfAllDirsSmallerThan(100000))
}

func (d *Day07) Part2() string {

	//res := int64(0)

	total := 70000000
	required := 30000000

	free := total - d.dir.Size()

	deleteMin := required - free
	dirs := d.dir.DirsGreaterThan(deleteMin)

	dirmap := make(map[string]int, len(dirs))
	keys := make([]string, 0, len(dirs))
	for _, dir := range dirs {
		dirmap[dir.Name] = dir.Size()
		keys = append(keys, dir.Name)
	}

	sort.SliceStable(keys, func(i, j int) bool {
		return dirmap[keys[i]] < dirmap[keys[j]]
	})

	return fmt.Sprintf("%d", dirmap[keys[0]])
}
