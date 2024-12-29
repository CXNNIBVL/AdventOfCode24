package main

import (
	"bufio"
	"os"
)

const (
	FILE      = "../inputs/D08/input"
	TEST_FILE = "../inputs/D08/input_test"
)

func readAllLines(f *os.File) []string {
	lines := []string{}

	sc := bufio.NewScanner(f)

	for sc.Scan() {
		lines = append(lines, sc.Text())
	}

	return lines
}

type Frequency rune
type Antenna struct {
	x, y int
	freq Frequency
}

const (
	OBJ_NOTHING rune = '.'
)

func parseAntennas(lines []string) []Antenna {
	antennas := []Antenna{}

	for row, line := range lines {
		rline := []rune(line)

		for col, f := range rline {
			if f == OBJ_NOTHING {
				continue
			}

			antennas = append(antennas, Antenna{x: col, y: row, freq: Frequency(f)})
		}
	}

	return antennas
}

type AntiNode struct {
	x, y  int
	freqs []Frequency
}

func calcAllAntiNodes(xbound, ybound int, antennas []Antenna) []AntiNode {
	// TODO:
	return nil
}

func main() {
	f, err := os.Open(TEST_FILE)

	if err != nil {
		panic(err)
	}

	defer f.Close()

	lines := readAllLines(f)
	antennas := parseAntennas(lines)

	xbound, ybound := len(lines[0]), len(lines)

	antiNodes := calcAllAntiNodes(xbound, ybound, antennas)
}
