package challenges

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/jcheyer/aoc2021/lib"
)

type Day04 struct {
	rawInput []string
}

func (d *Day04) Load(file string) error {

	if len(d.rawInput) > 0 {
		return nil
	}
	var err error
	d.rawInput, err = lib.ReadInputLines(file)
	if err != nil {
		return err
	}

	return nil
}

func (d *Day04) Part1() string {

	res := int64(0)
	for _, line := range d.rawInput {
		s1, s2 := line2sections(line)
		if fullOverlap(s1, s2) {
			res++
		}
	}

	return fmt.Sprintf("%d", res)
}

func (d *Day04) Part2() string {

	res := int64(0)

	for _, line := range d.rawInput {
		s1, s2 := line2sections(line)
		if overlap(s1, s2) {
			res++
		}
	}

	return fmt.Sprintf("%d", res)
}

type section struct {
	start int64
	end   int64
}

func line2section(line string) section {
	bounderies := strings.Split(line, "-")
	s := section{}
	s.start, _ = strconv.ParseInt(bounderies[0], 10, 64)
	s.end, _ = strconv.ParseInt(bounderies[1], 10, 64)

	return s
}

func line2sections(line string) (s1, s2 section) {
	secs := strings.Split(line, ",")
	s1 = line2section(secs[0])
	s2 = line2section(secs[1])

	return s1, s2
}

func fullOverlap(s1, s2 section) bool {
	if s1.start >= s2.start && s1.end <= s2.end ||
		s2.start >= s1.start && s2.end <= s1.end {

		return true
	}

	return false
}

func overlap(s1, s2 section) bool {
	if s1.end < s2.start || s2.end < s1.start {
		return false
	}
	return true
}
