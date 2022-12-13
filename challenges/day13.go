package challenges

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"reflect"
	"sort"
	"strings"
)

type d13pair struct {
	p1, p2 string
	x1, x2 any
}
type Day13 struct {
	data []d13pair
}

func (d *Day13) Load(file string) error {

	data, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}
	for _, pairs := range strings.Split(string(data), "\n\n") {
		parts := strings.Split(pairs, "\n")

		x1, err := d.s2a(parts[0])
		if err != nil {
			return err
		}
		x2, err := d.s2a(parts[1])
		if err != nil {
			return err
		}

		d.data = append(d.data,
			d13pair{
				p1: parts[0],
				p2: parts[1],
				x1: x1,
				x2: x2})

	}
	//spew.Dump(d.data)
	return nil
}
func (d *Day13) s2a(s string) (res any, err error) {

	err = json.Unmarshal([]byte(s), &res)
	return res, err
}

func (d *Day13) Part1() string {
	res := 0

	for i, pair := range d.data {
		if d.solve(pair.x1, pair.x2) <= 0 {
			res += i + 1
		}
	}

	return fmt.Sprintf("%d", res)
}

func (d *Day13) Part2() string {
	packets := []any{}
	for _, packet := range d.data {
		packets = append(packets, packet.x1, packet.x2)
	}

	var div1 any
	var div2 any
	_ = json.Unmarshal([]byte("[[2]]"), &div1)
	_ = json.Unmarshal([]byte("[[6]]"), &div2)

	packets = append(packets, div1, div2)

	sort.Slice(packets, func(i, j int) bool {
		return d.solve(packets[i], packets[j]) < 0
	})

	res := 1
	for k, v := range packets {
		mVal, _ := json.Marshal(v)
		if string(mVal) == "[[2]]" || string(mVal) == "[[6]]" {
			res *= k + 1
		}
	}

	return fmt.Sprintf("%d", res)
}

func (d Day13) solve(x1, x2 any) int {

	// both numbers
	v1, ok1 := x1.(float64)
	v2, ok2 := x2.(float64)
	if ok1 && ok2 {
		return int(v1) - int(v2)
	}

	var l1 []any
	if reflect.TypeOf(x1).String() == "float64" {
		l1 = []any{x1}
	} else {
		l1 = x1.([]any)
	}

	var l2 []any
	if reflect.TypeOf(x2).String() == "float64" {
		l2 = []any{x2}
	} else {
		l2 = x2.([]any)
	}

	for i := range l1 {
		// right side out of items so +1 (wrong order)
		if len(l2) <= i {
			return 1
		}
		if r := d.solve(l1[i], l2[i]); r != 0 {
			return r
		}
	}
	if len(l1) == len(l2) {
		return 0
	}

	return -1

}
