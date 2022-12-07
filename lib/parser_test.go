package lib

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadInputLines(t *testing.T) {
	expected := []string{"157", "148", "149", "146", "144", "145", "162", "163", "164", "166"}
	s1, err := ReadInputLines("testfiles/input.txt")
	assert.NoError(t, err)
	assert.Len(t, s1, 10)
	assert.Equal(t, expected, s1)

	_, err = ReadInputLines("not found.txt")
	assert.Error(t, err)
}

func TestParseLines2Int64(t *testing.T) {
	input := []string{"157", "148", "149", "146", "144", "145", "162", "163", "164", "166"}
	expected := []int64{157, 148, 149, 146, 144, 145, 162, 163, 164, 166}
	output, err := ParseLines2Int(input)
	assert.NoError(t, err)
	assert.Equal(t, expected, output)

	_, err = ParseLines2Int([]string{"not a number", "123"})
	assert.Error(t, err)
}

func TestParseLines2Nav1(t *testing.T) {
	s1, err := ReadInputLines("testfiles/d2demo.txt")
	assert.NoError(t, err)
	hor, depth := ParseLines2Nav1(s1)
	assert.Equal(t, 10, depth)
	assert.Equal(t, 15, hor)
}

func TestParseLines2Nav2(t *testing.T) {
	s1, err := ReadInputLines("testfiles/d2demo.txt")
	assert.NoError(t, err)
	hor, depth := ParseLines2Nav2(s1)
	assert.Equal(t, 60, depth)
	assert.Equal(t, 15, hor)
}

func TestParseAndAggregate(t *testing.T) {
	s1, err := ReadInputLinesWithEmptyLines("testfiles/aggregate.txt")
	assert.NoError(t, err)
	agg, err := ParseAndAggregateLines2Int(s1)
	assert.NoError(t, err)
	assert.Len(t, agg, 5)
	assert.Equal(t, agg[0], int64(6000))
	assert.Equal(t, agg[4], int64(10000))

}
