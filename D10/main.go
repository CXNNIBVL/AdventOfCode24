package main

import (
	"bufio"
	"fmt"
	"os"
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

func isInsideBounds(x, y, xbound, ybound int) bool {
	return x < xbound && x >= 0 && y < ybound && y >= 0
}

const (
	ACTION_GO_LEFT = iota
	ACTION_GO_RIGHT
	ACTION_GO_UP
	ACTION_GO_DOWN
	ACTION_GO_XX
)

func walkTrail(xstart, ystart int, m [][]int, prevDir, xbound, ybound int) int {

	score := 0

	current := m[ystart][xstart]

	if current == 9 {
		score = score + 1
	}

	// Go right
	if newx, newy := xstart+1, ystart; prevDir != ACTION_GO_LEFT && isInsideBounds(newx, newy, xbound, ybound) {
		next := m[newy][newx]
		if next == current+1 {
			score = score + walkTrail(newx, newy, m, ACTION_GO_RIGHT, xbound, ybound)
		}
	}

	// Go left
	if newx, newy := xstart-1, ystart; prevDir != ACTION_GO_RIGHT && isInsideBounds(newx, newy, xbound, ybound) {
		next := m[newy][newx]
		if next == current+1 {
			score = score + walkTrail(newx, newy, m, ACTION_GO_LEFT, xbound, ybound)
		}
	}

	// Go up
	if newx, newy := xstart, ystart-1; prevDir != ACTION_GO_DOWN && isInsideBounds(newx, newy, xbound, ybound) {
		next := m[newy][newx]
		if next == current+1 {
			score = score + walkTrail(newx, newy, m, ACTION_GO_UP, xbound, ybound)
		}
	}

	// Go down
	if newx, newy := xstart, ystart+1; prevDir != ACTION_GO_UP && isInsideBounds(newx, newy, xbound, ybound) {
		next := m[newy][newx]
		if next == current+1 {
			score = score + walkTrail(newx, newy, m, ACTION_GO_DOWN, xbound, ybound)
		}
	}

	return score
}

func getTrailHeadScore(row, col int, m [][]int, ch chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	ch <- walkTrail(col, row, m, ACTION_GO_XX, len(m[0]), len(m))
}

func main() {
	f, err := os.Open(TEST_FILE)

	if err != nil {
		panic(err)
	}

	defer f.Close()

	m := parseMap(f)

	scores := 0

	// ch := make(chan int)
	// var wg sync.WaitGroup

	for row, line := range m {
		for col, c := range line {
			if c == 0 {
				// wg.Add(1)
				// go getTrailHeadScore(row, col, m, ch, &wg)
				// TODO: try syncronous first
				scores = scores + walkTrail(col, row, m, ACTION_GO_XX, len(m[0]), len(m))
			}
		}
	}

	// go func() {
	// 	wg.Wait()
	// 	close(ch)
	// }()

	// for s := range ch {
	// 	scores = scores + s
	// }

	fmt.Println(scores)
}
