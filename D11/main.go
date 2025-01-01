package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/CXNNIBVL/goutil/iter"
)

const (
	FILE       = "../inputs/D11/input"
	TEST_FILE  = "../inputs/D11/input_test"
	TEST_FILE2 = "../inputs/D11/input_test2"
)

type Stone = int

func parseStones(f *os.File) []Stone {
	stones := []Stone{}

	sc := bufio.NewScanner(f)

	if sc.Scan() {
		line := sc.Text()

		numberStrs := strings.Split(line, " ")

		for _, numStr := range numberStrs {
			num, _ := strconv.ParseInt(numStr, 10, 64)
			stones = append(stones, Stone(num))
		}
	}

	return stones
}

func applyBlinkV2(stone Stone) []Stone {
	stones := []Stone{}

	numDigits := 0
	for s := stone; s != 0; s = s / 10 {
		numDigits++
	}

	if stone == 0 {
		stones = append(stones, Stone(1))
	} else if numDigits%2 == 0 {
		hi := stone / int(math.Pow10(numDigits/2))
		lo := stone - hi*int(math.Pow10(numDigits/2))
		stones = append(stones, Stone(hi), Stone(lo))
	} else {
		stones = append(stones, Stone(stone*2024))
	}

	return stones
}

type StoneCounter = map[Stone]int

func main() {
	f, err := os.Open(FILE)

	if err != nil {
		panic(err)
	}

	defer f.Close()

	stones := parseStones(f)
	numBlinks := 75

	m := make(map[int]StoneCounter)
	m[0] = make(StoneCounter)
	for _, stone := range stones {
		m[0][stone] = m[0][stone] + 1
	}

	for i := range iter.Interval(1, numBlinks+1) {
		m[i] = make(StoneCounter)

		for stone, numStones := range m[i-1] {
			for _, stone := range applyBlinkV2(stone) {
				m[i][stone] = m[i][stone] + 1*numStones
			}
		}

		if i == 25 || i == numBlinks {
			items := 0
			for _, v := range m[i] {
				items = items + v
			}
			fmt.Printf("Blink %d: Items = %d\n", i, items)
		}
	}
}
