package challenges

import (
	"fmt"
	"sort"

	"github.com/davecgh/go-spew/spew"
	"github.com/jcheyer/aoc2021/challenges/internal/monkeys"
	"github.com/jcheyer/aoc2021/lib"
)

type Day11 struct {
	rawInput []string
	monkeys  map[int]*monkeys.Monkey
	monkeys2 map[int]*monkeys.Monkey

	zauberzahl int64
}

func (d *Day11) Load(file string) error {

	if len(d.rawInput) > 0 {
		return nil
	}
	var err error
	d.rawInput, err = lib.ReadInputLinesRaw(file)
	if err != nil {
		return err
	}
	d.monkeys = monkeys.New(d.rawInput)
	d.monkeys2 = monkeys.New(d.rawInput)
	d.zauberzahl = 1
	for _, m := range d.monkeys {
		d.zauberzahl *= m.Test
	}

	//spew.Dump(d.monkeys)
	return nil
}

func (d *Day11) Part1() string {

	d.monkeys = d.Do(1, 20, func(x int64) int64 { return x / 3 })

	tmp := []int{}
	for m := 0; m < len(d.monkeys); m++ {
		tmp = append(tmp, d.monkeys[m].Counter)
		//spew.Dump(d.monkeys[m].Counter)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(tmp)))

	return fmt.Sprintf("%d", tmp[0]*tmp[1])
}

func (d *Day11) Part2() string {

	d.monkeys = d.Do(2, 1, func(x int64) int64 { return x % d.zauberzahl })
	d.DumpCounters()
	d.monkeys = d.Do(2, 19, func(x int64) int64 { return x % d.zauberzahl })
	d.DumpCounters()
	d.monkeys = d.Do(2, 980, func(x int64) int64 { return x % d.zauberzahl })
	d.DumpCounters()
	d.monkeys = d.Do(2, 9000, func(x int64) int64 { return x % d.zauberzahl })
	d.DumpCounters()

	tmp := []int{}
	fmt.Println("---------------------")
	for m := 0; m < len(d.monkeys2); m++ {
		tmp = append(tmp, d.monkeys2[m].Counter)

		spew.Dump(d.monkeys2[m].Counter)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(tmp)))

	return fmt.Sprintf("%d", int64(tmp[0]*tmp[1]))
}

func (d *Day11) Do(what int, count int, f func(int64) int64) map[int]*monkeys.Monkey {
	monkeys := map[int]*monkeys.Monkey{}
	if what == 1 {
		monkeys = d.monkeys
	} else {
		monkeys = d.monkeys2
	}

	for i := 0; i < count; i++ {
		for m := 0; m < len(monkeys); m++ {
			for _, item := range monkeys[m].Items {
				monkeys[m].Counter++
				val := monkeys[m].Operation(item)
				val2 := f(val)

				if val2%monkeys[m].Test == 0 {
					monkeys[monkeys[m].TrueTarget].Items = append(monkeys[monkeys[m].TrueTarget].Items, val2)
				} else {
					monkeys[monkeys[m].FalseTarget].Items = append(monkeys[monkeys[m].FalseTarget].Items, val2)
				}
			}
			monkeys[m].Items = []int64{}
		}
	}
	return monkeys
}

func (d *Day11) DumpCounters() {
	fmt.Println("------- Dump ------------")
	for m := 0; m < len(d.monkeys2); m++ {
		spew.Dump(d.monkeys2[m].Counter)
	}
}
