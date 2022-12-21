package main

import (
	"fmt"
	"log"
	"os"

	"github.com/jcheyer/aoc2021/challenges"
)

func main() {
	if len(os.Args) < 2 {
		panic("Please provide the problem name (e.g. 'day01') as parameter.")
	}

	challengeName := os.Args[1]

	challengeMap := make(map[string]challenges.Challenge)

	challengeMap["day01"] = &challenges.Day01{}
	challengeMap["day02"] = &challenges.Day02{}
	challengeMap["day03"] = &challenges.Day03{}
	challengeMap["day04"] = &challenges.Day04{}
	challengeMap["day05"] = &challenges.Day05{}
	challengeMap["day06"] = &challenges.Day06{}
	challengeMap["day07"] = &challenges.Day07{}
	challengeMap["day08"] = &challenges.Day08{}
	challengeMap["day09"] = &challenges.Day09{}
	challengeMap["day10"] = &challenges.Day10{}
	challengeMap["day11"] = &challenges.Day11{}
	challengeMap["day12"] = &challenges.Day12{}
	challengeMap["day13"] = &challenges.Day13{}
	challengeMap["day14"] = &challenges.Day14{}
	challengeMap["day15"] = &challenges.Day15{}
	challengeMap["day16"] = &challenges.Day16{}
	challengeMap["day17"] = &challenges.Skipped{}
	challengeMap["day18"] = &challenges.Day18{}
	challengeMap["day19"] = &challenges.Skipped{}

	challengeMap["day20"] = &challenges.Skipped{}
	challengeMap["day21"] = &challenges.Day21{}
	challengeMap["day22"] = &challenges.Skipped{}
	challengeMap["day23"] = &challenges.Skipped{}
	challengeMap["day24"] = &challenges.Skipped{}

	problem, defined := challengeMap[challengeName]
	if !defined {
		log.Panicf("challenge not found: %s\n", challengeName)
	}

	if err := problem.Load("input/" + challengeName + ".txt"); err != nil {
		log.Panicf("load error(%s): %+v\n", challengeName, err)
	}

	fmt.Printf("Part1: %s\n", problem.Part1())
	fmt.Printf("Part2: %s\n", problem.Part2())

}
