package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"sync"
)

const (
	FILE      = "../inputs/D10/input"
	TEST_FILE = "../inputs/D10/input_test"
)

func parseMap(f *os.File) [][]int {
	sc := bufio.NewScanner(f)

	m := [][]int{}
	for sc.Scan() {
		s := sc.Text()

		l := []int{}
		for _, c := range s {
			v, _ := strconv.ParseInt(string(c), 10, 64)
			l = append(l, int(v))
		}
		m = append(m, l)
	}

	return m
}

type Vec2 struct {
	x, y int
}

func isInsideBounds(p Vec2, xbound, ybound int) bool {
	return p.x < xbound && p.x >= 0 && p.y < ybound && p.y >= 0
}

const (
	DIR_GO_LEFT = iota
	DIR_GO_RIGHT
	DIR_GO_UP
	DIR_GO_DOWN
	DIR_GO_XX
)

func walkTrail(start Vec2, m [][]int, prevDir, xbound, ybound int, alreadyFound []Vec2) []Vec2 {

	current := m[start.y][start.x]

	if current == 9 {
		containsCurrent := slices.ContainsFunc(alreadyFound, func(v Vec2) bool {
			return v.x == start.x && v.y == start.y
		})

		if containsCurrent {
			return alreadyFound
		}

		return append(alreadyFound, start)
	}

	// Go right
	newStart := Vec2{x: start.x + 1, y: start.y}
	if prevDir != DIR_GO_LEFT && isInsideBounds(newStart, xbound, ybound) {
		next := m[newStart.y][newStart.x]
		if next == current+1 {
			alreadyFound = walkTrail(newStart, m, DIR_GO_RIGHT, xbound, ybound, alreadyFound)
		}
	}

	// Go left
	newStart = Vec2{x: start.x - 1, y: start.y}
	if prevDir != DIR_GO_RIGHT && isInsideBounds(newStart, xbound, ybound) {
		next := m[newStart.y][newStart.x]
		if next == current+1 {
			alreadyFound = walkTrail(newStart, m, DIR_GO_LEFT, xbound, ybound, alreadyFound)
		}
	}

	// Go up
	newStart = Vec2{x: start.x, y: start.y - 1}
	if prevDir != DIR_GO_DOWN && isInsideBounds(newStart, xbound, ybound) {
		next := m[newStart.y][newStart.x]
		if next == current+1 {
			alreadyFound = walkTrail(newStart, m, DIR_GO_UP, xbound, ybound, alreadyFound)
		}
	}

	// Go down
	newStart = Vec2{x: start.x, y: start.y + 1}
	if prevDir != DIR_GO_UP && isInsideBounds(newStart, xbound, ybound) {
		next := m[newStart.y][newStart.x]
		if next == current+1 {
			alreadyFound = walkTrail(newStart, m, DIR_GO_DOWN, xbound, ybound, alreadyFound)
		}
	}

	return alreadyFound
}

func walkTrailParallel(start Vec2, m [][]int, ch chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	ch <- len(walkTrail(start, m, DIR_GO_XX, len(m[0]), len(m), []Vec2{}))
}

func main() {
	f, err := os.Open(FILE)

	if err != nil {
		panic(err)
	}

	defer f.Close()

	m := parseMap(f)

	scores := 0

	ch := make(chan int)
	var wg sync.WaitGroup

	for row, line := range m {
		for col, c := range line {
			if c == 0 {
				wg.Add(1)
				start := Vec2{x: col, y: row}
				go walkTrailParallel(start, m, ch, &wg)
			}
		}
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	for s := range ch {
		scores = scores + s
	}

	fmt.Println("Part1: ", scores)
}
