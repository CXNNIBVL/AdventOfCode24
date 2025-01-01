package main

import (
	"bufio"
	"os"
)

const (
	FILE       = "../inputs/D12/input"
	TEST_FILE1 = "../inputs/D12/input_test1"
	TEST_FILE2 = "../inputs/D12/input_test2"
	TEST_FILE3 = "../inputs/D12/input_test3"
)

type Vec2 struct {
	x, y int
}

func parseMatrix(f *os.File) []string {
	sc := bufio.NewScanner(f)

	lines := []string{}

	for sc.Scan() {
		lines = append(lines, sc.Text())
	}

	return lines
}

func findAreaAndPerimeter(start Vec2, mat []string) (area, perim int) {
	area = 0
	perim = 0

	return
}

func main() {
	f, err := os.Open(TEST_FILE1)

	if err != nil {
		panic(err)
	}

	defer f.Close()

	mat := parseMatrix(f)

	price := 0

	for row, line := range mat {
		for col, r := range line {
			if r != '.' {
				area, perim := findAreaAndPerimeter(Vec2{x: col, y: row}, mat)
				price = price + area*perim
			}
		}
	}

}
