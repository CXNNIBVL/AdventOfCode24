package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

const (
	FILE       = "../inputs/D11/input"
	TEST_FILE  = "../inputs/D11/input_test"
	TEST_FILE2 = "../inputs/D11/input_test2"
)

type Stone struct {
	number int64
}

func parseStones(f *os.File) []Stone {
	stones := []Stone{}

	sc := bufio.NewScanner(f)

	if sc.Scan() {
		line := sc.Text()

		numberStrs := strings.Split(line, " ")

		for _, numStr := range numberStrs {
			num, _ := strconv.ParseInt(numStr, 10, 64)
			stones = append(stones, Stone{number: num})
		}
	}

	return stones
}

func applyBlinks(stone Stone, numBlinks int) []Stone {
	stones := []Stone{}

	if numBlinks == 0 {
		return []Stone{stone}
	}

	numDigits := 0
	for s := stone.number; s != 0; s = s / 10 {
		numDigits++
	}

	if stone.number == 0 {
		stones = append(stones, Stone{number: 1})
	} else if numDigits%2 == 0 {
		hi := stone.number / int64(math.Pow10(numDigits/2))
		lo := stone.number - hi*int64(math.Pow10(numDigits/2))
		stones = append(stones, Stone{number: hi}, Stone{number: lo})
	} else {
		stones = append(stones, Stone{number: stone.number * 2024})
	}

	newStones := []Stone{}
	for _, st := range stones {
		newStones = append(newStones, applyBlinks(st, numBlinks-1)...)
	}

	return newStones
}

func main() {
	f, err := os.Open(FILE)

	if err != nil {
		panic(err)
	}

	defer f.Close()

	stones := parseStones(f)
	numBlinks := 25

	newStones := []Stone{}
	for _, stone := range stones {
		newStones = append(newStones, applyBlinks(stone, numBlinks)...)
	}

	fmt.Println("Part1: ", len(newStones))
}
