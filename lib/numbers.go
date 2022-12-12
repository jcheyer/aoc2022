package lib

import "math"

const MaxUint = ^uint(0)
const MinUint = 0
const MaxInt = int(MaxUint >> 1)
const MinInt = -MaxInt - 1

func HighInt(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func LowInt(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func HighUInt64(a uint64, b uint64) uint64 {
	if a > b {
		return a
	}
	return b
}

func LowUInt64(a uint64, b uint64) uint64 {
	if a < b {
		return a
	}
	return b
}

func AbsDiffUInt64(a, b uint64) uint64 {
	if a > b {
		return a - b
	}

	return b - a
}

func HighInt64(a int64, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

func LowInt64(a int64, b int64) int64 {
	if a < b {
		return a
	}
	return b
}

func OnOff(a, b int) bool {
	return math.Abs(float64(a-b)) <= 1
}
