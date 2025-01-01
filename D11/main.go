package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"sync"

	"github.com/CXNNIBVL/goutil/iter"
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

	// newStones := []Stone{}
	// for _, st := range stones {
	// 	newStones = append(newStones, applyBlinks(st, numBlinks-1)...)
	// }

	return multiplexStones(stones, numBlinks-1)
}

func applyBlinksConcurrently(stone Stone, blinks int, ch chan<- []Stone, wg *sync.WaitGroup) {
	defer wg.Done()
	ch <- applyBlinks(stone, blinks)
}

func multiplexStones(stones []Stone, blinks int) []Stone {
	ch := make(chan []Stone)
	var wg sync.WaitGroup

	newStones := []Stone{}

	for _, stone := range stones {
		wg.Add(1)
		go applyBlinksConcurrently(stone, blinks, ch, &wg)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	for s := range ch {
		newStones = append(newStones, s...)
	}

	return newStones
}

func applyBlinkV2(stone Stone) []Stone {
	stones := []Stone{}

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

	return stones
}

func applyBlinkConcurrentlyV2(stones []Stone, thresh, maxGoroutines int) []Stone {
	chunks := [][]Stone{}

	tmp := stones
	for len(tmp) > 0 {
		end := thresh

		if l := len(tmp); l < thresh {
			end = l
		}

		chunk := tmp[:end]
		chunks = append(chunks, chunk)
		tmp = tmp[end:]
	}

	var wg sync.WaitGroup
	ch := make(chan []Stone)

	ctr := 0
	for _, chunk := range chunks {
		ctr++
		wg.Add(1)
		go func() {
			defer wg.Done()
			n := []Stone{}
			for _, stone := range chunk {
				n = append(n, applyBlinkV2(stone)...)
			}
			ch <- n
		}()

		if ctr > maxGoroutines {
			fmt.Println("Num Goroutines: ", ctr)
			wg.Wait()
			ctr = 0
		}
	}

	go func() {
		fmt.Println("Num Goroutines: ", ctr)
		wg.Wait()
		close(ch)
	}()

	newStones := []Stone{}
	for list := range ch {
		newStones = append(newStones, list...)
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
	numBlinks := 75
	maxGoroutines := 800

	newStones := stones
	thresh := 10
	for i := range iter.Interval(1, numBlinks+1) {
		fmt.Printf("Blink: %d, Len: %d, Thresh: %d\n", i, len(newStones), thresh)
		newStones = applyBlinkConcurrentlyV2(newStones, thresh, maxGoroutines)
		thresh = int(float64(thresh) * 1.25)

		if i == 24 {
			fmt.Println("Part1: ", len(newStones))
		}
	}

	fmt.Println("Part2: ", len(newStones))
}
