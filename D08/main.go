package main

import (
	"bufio"
	"fmt"
	"os"
	"sync"

	"github.com/CXNNIBVL/goutil/math"
)

const (
	FILE       = "../inputs/D08/input"
	TEST_FILE  = "../inputs/D08/input_test"
	TEST_FILE1 = "../inputs/D08/input_test1"
	TEST_FILE2 = "../inputs/D08/input_test2"
	TEST_FILE3 = "../inputs/D08/input_test3"
	EN_DBG     = false
)

func readAllLines(f *os.File) []string {
	lines := []string{}

	sc := bufio.NewScanner(f)

	for sc.Scan() {
		lines = append(lines, sc.Text())
	}

	return lines
}

type Vec2 struct {
	x, y int
}

type Frequency rune
type Antenna struct {
	pos  Vec2
	freq Frequency
}

const (
	OBJ_NOTHING rune = '.'
)

func parseAntennas(lines []string) map[Frequency][]Vec2 {
	m := make(map[Frequency][]Vec2)

	for row, line := range lines {
		rline := []rune(line)

		for col, r := range rline {

			if r == OBJ_NOTHING {
				continue
			}

			f := Frequency(r)

			if m[f] == nil {
				m[f] = append([]Vec2{}, Vec2{x: col, y: row})
			} else {
				m[f] = append(m[f], Vec2{x: col, y: row})
			}
		}
	}

	return m
}

type AntiNodeResult struct {
	freq      Frequency
	positions []Vec2
}

func isPointInBound(v Vec2, xbound, ybound int) bool {
	return v.x < xbound && v.x >= 0 && v.y < ybound && v.y >= 0
}

func findAntiNodePositions(freq Frequency, xbound, ybound int, antennaPoints []Vec2, ch chan<- AntiNodeResult, wg *sync.WaitGroup) {
	defer wg.Done()

	foundPositions := []Vec2{}

	points := antennaPoints
	for len(points) > 1 {
		for ix := 1; ix < len(points); ix++ {
			p1, p2 := points[0], points[ix]

			if p1.x == p2.x && p1.y == p2.y {
				continue
			}

			hdiff, vdiff := math.Abs(p1.x-p2.x), math.Abs(p1.y-p2.y)
			anti1, anti2 := Vec2{x: -1, y: -1}, Vec2{x: -1, y: -1}
			hflip, vflip := 1, 1

			// Check if p1 is rigth of / below p2
			if p1.x > p2.x || p1.y > p2.y {
				// Flip addition direction
				hflip = -1
			}

			if p1.y > p2.y {
				// Flip addition direction
				vflip = -1
			}

			// Note: The diagonal doesn't have to be a perfect diagonal
			// like you would find going thru corner to corner of a square
			anti1.y, anti2.y = p1.y-vflip*vdiff, p2.y+vflip*vdiff
			anti1.x, anti2.x = p1.x-hflip*hdiff, p2.x+hflip*hdiff

			if isPointInBound(anti1, xbound, ybound) {
				foundPositions = append(foundPositions, anti1)
			}

			if isPointInBound(anti2, xbound, ybound) {
				foundPositions = append(foundPositions, anti2)
			}
		}

		points = points[1:]
	}

	ch <- AntiNodeResult{positions: foundPositions, freq: freq}
}

func main() {
	f, err := os.Open(FILE)

	if err != nil {
		panic(err)
	}

	defer f.Close()

	lines := readAllLines(f)
	antennaMap := parseAntennas(lines)

	xbound, ybound := len(lines[0]), len(lines)

	if EN_DBG {
		fmt.Println("### Antennas")
		for k, va := range antennaMap {
			fmt.Printf("--- Freq: %s ---\n", string(k))
			for _, v := range va {
				fmt.Printf("%+v\n", v)
			}
		}
	}

	ch := make(chan AntiNodeResult)
	var wg sync.WaitGroup

	for freq, antennaPoints := range antennaMap {
		wg.Add(1)
		go findAntiNodePositions(freq, xbound, ybound, antennaPoints, ch, &wg)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	antiNodeMap := make(map[Vec2][]Frequency)

	for r := range ch {
		for _, p := range r.positions {
			antiNodeMap[p] = append(antiNodeMap[p], r.freq)
		}
	}

	if EN_DBG {
		fmt.Println("### AntiNodes")
		for p, freqs := range antiNodeMap {
			sfreqs := string(freqs[0])
			freqs = freqs[1:]

			for _, r := range freqs {
				sfreqs = sfreqs + ", " + string(r)
			}

			fmt.Printf("p: %+v, f: [%s]\n", p, sfreqs)
		}
	}

	fmt.Println(len(antiNodeMap))
}
