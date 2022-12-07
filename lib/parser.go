package lib

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

// ReadString reads a file into a slice of lines
func ReadInputLines(filename string) ([]string, error) {
	maxCapacity := 65536

	file, err := os.Open(filename)
	if err != nil {
		return []string{}, err
	}
	defer file.Close()
	result := make([]string, 0)

	scanner := bufio.NewScanner(file)
	buf := make([]byte, maxCapacity)
	scanner.Buffer(buf, maxCapacity)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if len(line) > 0 {
			result = append(result, line)
		}
	}

	return result, nil
}

// ReadString reads a file into a slice of lines
func ReadInputLinesRaw(filename string) ([]string, error) {
	maxCapacity := 65536

	file, err := os.Open(filename)
	if err != nil {
		return []string{}, err
	}
	defer file.Close()
	result := make([]string, 0)

	scanner := bufio.NewScanner(file)
	buf := make([]byte, maxCapacity)
	scanner.Buffer(buf, maxCapacity)

	for scanner.Scan() {
		line := scanner.Text()

		result = append(result, line)

	}

	return result, nil
}

// ReadString reads a file into a slice of lines
func ReadInputLinesWithEmptyLines(filename string) ([]string, error) {
	maxCapacity := 65536

	file, err := os.Open(filename)
	if err != nil {
		return []string{}, err
	}
	defer file.Close()
	result := make([]string, 0)

	scanner := bufio.NewScanner(file)
	buf := make([]byte, maxCapacity)
	scanner.Buffer(buf, maxCapacity)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		result = append(result, line)
	}

	return result, nil
}

// ParseLines2Int64 takes a slice of strings and returns a slice of int64
func ParseLines2Int(lines []string) ([]int64, error) {
	res := make([]int64, len(lines))
	for i, line := range lines {
		convValue, err := strconv.ParseInt(line, 10, 64)
		if err != nil {
			return nil, err
		}
		res[i] = convValue
	}
	return res, nil
}

func ParseAndAggregateLines2Int(lines []string) (map[int]int64, error) {
	res := make(map[int]int64, 1)
	currentIndex := 0
	res[currentIndex] = 0
	for _, line := range lines {
		if line == "" {
			currentIndex += 1
			continue
		}
		convValue, err := strconv.ParseInt(line, 10, 64)
		if err != nil {
			return nil, err
		}
		res[currentIndex] += convValue
	}

	return res, nil
}

func ParseLines2Nav1(lines []string) (int, int) {
	hor := 0
	depth := 0
	for _, line := range lines {
		parts := strings.Fields(line)
		pval, _ := strconv.Atoi(parts[1])
		switch parts[0] {
		case "up":
			depth -= pval
		case "down":
			depth += pval
		case "forward":
			hor += pval
		}
	}
	return hor, depth
}

func ParseLines2Nav2(lines []string) (int, int) {
	aim := 0
	hor := 0
	depth := 0
	for _, line := range lines {
		parts := strings.Fields(line)
		pval, _ := strconv.Atoi(parts[1])
		switch parts[0] {
		case "up":
			aim -= pval
		case "down":
			aim += pval
		case "forward":
			hor += pval
			depth += pval * aim
		}
	}
	return hor, depth
}
