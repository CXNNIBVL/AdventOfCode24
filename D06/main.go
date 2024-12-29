package main

import (
	"bufio"
	"fmt"
	"os"
	"sync"

	"github.com/CXNNIBVL/goutil/iter"
)

const (
	FILE      = "../inputs/D06/input"
	FILE_TEST = "../inputs/D06/input_test"
)

const (
	OBJ_OBSTACLE    rune = '#'
	OBJ_NOTHING     rune = '.'
	OBJ_OOB         rune = '\x00'
	OBJ_GUARD_LEFT  rune = '<'
	OBJ_GUARD_RIGHT rune = '>'
	OBJ_GUARD_UP    rune = '^'
	OBJ_GUARD_DOWN  rune = 'v'
)

func getFileLines(f *os.File) []string {
	lines := []string{}

	sc := bufio.NewScanner(f)

	for sc.Scan() {
		lines = append(lines, sc.Text())
	}

	return lines
}

type Mat2 [][]rune

func (m *Mat2) NumRowsCols() (r, c int) {
	return len((*m)), len((*m)[0])
}

func runeIsGuard(r rune) bool {
	return r == OBJ_GUARD_DOWN || r == OBJ_GUARD_UP || r == OBJ_GUARD_LEFT || r == OBJ_GUARD_RIGHT
}

type Vec2 struct {
	x, y int
}

type Guard struct {
	position  Vec2
	direction rune
}

func (g *Guard) advance() {
	if g.direction == OBJ_GUARD_DOWN {
		g.position.y = g.position.y + 1
		return
	}

	if g.direction == OBJ_GUARD_UP {
		g.position.y = g.position.y - 1
		return
	}

	if g.direction == OBJ_GUARD_LEFT {
		g.position.x = g.position.x - 1
		return
	}

	if g.direction == OBJ_GUARD_RIGHT {
		g.position.x = g.position.x + 1
		return
	}
}

func (g *Guard) turn() {
	if g.direction == OBJ_GUARD_DOWN {
		g.direction = OBJ_GUARD_LEFT
		return
	}

	if g.direction == OBJ_GUARD_UP {
		g.direction = OBJ_GUARD_RIGHT
		return
	}

	if g.direction == OBJ_GUARD_LEFT {
		g.direction = OBJ_GUARD_UP
		return
	}

	if g.direction == OBJ_GUARD_RIGHT {
		g.direction = OBJ_GUARD_DOWN
		return
	}
}

func inspectMapIndex(y, x int, m Mat2) rune {
	ybound, xbound := m.NumRowsCols()

	if x >= xbound || x < 0 || y >= ybound || y < 0 {
		return OBJ_OOB
	}

	return m[y][x]
}

type AdvanceResult int

const (
	ADV_OK  AdvanceResult = iota
	ADV_OOB AdvanceResult = iota
	ADV_OBS AdvanceResult = iota
)

func (g *Guard) canAdvance(m Mat2) AdvanceResult {

	var ahead rune

	if g.direction == OBJ_GUARD_RIGHT {
		ahead = inspectMapIndex(g.position.y, g.position.x+1, m)
	} else if g.direction == OBJ_GUARD_LEFT {
		ahead = inspectMapIndex(g.position.y, g.position.x-1, m)
	} else if g.direction == OBJ_GUARD_UP {
		ahead = inspectMapIndex(g.position.y-1, g.position.x, m)
	} else if g.direction == OBJ_GUARD_DOWN {
		ahead = inspectMapIndex(g.position.y+1, g.position.x, m)
	}

	if ahead == OBJ_OBSTACLE {
		return ADV_OBS
	}

	if ahead == OBJ_OOB {
		return ADV_OOB
	}

	return ADV_OK
}

func parseMapFromLines(lines []string) (m Mat2, g Guard) {

	m = Mat2{}

	for _, line := range lines {
		m = append(m, []rune(line))
	}

	nrows, ncols := m.NumRowsCols()

	// Find guard
	for row := range iter.Interval(0, nrows) {
		for col := range iter.Interval(0, ncols) {
			if r := m[row][col]; runeIsGuard(r) {
				g.position.x = col
				g.position.y = row
				g.direction = r
				return
			}
		}
	}

	panic("Did not find guard!")
}

type GameState struct {
	m  Mat2
	g  Guard
	hm map[Vec2]int
}

func (gs *GameState) loop(tickCb func(gs *GameState) bool) {
	for {
		// Heatmap
		gs.hm[gs.g.position]++

		if exit := tickCb(gs); exit {
			break
		}

		res := gs.g.canAdvance(gs.m)

		if res == ADV_OOB {
			break
		}

		if res == ADV_OBS {
			gs.g.turn()
			continue
		}

		gs.g.advance()
	}
}

func Part1(fileLines []string) {
	m, guard := parseMapFromLines(fileLines)
	gs := GameState{m: m, g: guard}

	gs.addHeatMap()
	gs.loop(func(gs *GameState) bool {
		return false
	})

	fmt.Println(len(gs.hm))
}

func (gs *GameState) addHeatMap() {
	gs.hm = make(map[Vec2]int)
}

func (gs *GameState) clearHeatMap() {
	gs.hm = make(map[Vec2]int)
}

func testAddingObstruction(wg *sync.WaitGroup, ch chan<- int, beginAtRow, numRows int, fileLines []string) {
	defer wg.Done()

	m, g := parseMapFromLines(fileLines)
	gs := GameState{m: m, g: g}
	gs.addHeatMap()

	_, ncols := gs.m.NumRowsCols()

	numGuardsStuck := 0

	for numRows > 0 {
		for col := range iter.Interval(0, ncols) {

			obj := inspectMapIndex(beginAtRow, col, gs.m)

			if obj == OBJ_OBSTACLE || (beginAtRow == gs.g.position.y && col == gs.g.position.x) {
				// skip
				continue
			}

			placedx, placedy := col, beginAtRow

			// Place obstacle
			gs.m[placedy][placedx] = OBJ_OBSTACLE

			stuck := false
			gs.loop(func(gs *GameState) bool {
				for _, v := range gs.hm {
					if v > 50 {
						stuck = true
						return true
					}
				}

				return false
			})

			if stuck {
				numGuardsStuck = numGuardsStuck + 1
				// fmt.Printf("%v\n", gs.hm)
			}

			// Cleanup placed obstacle
			gs.m[placedy][placedx] = obj
			// Clear heatmap
			gs.clearHeatMap()
			// Reset guard
			gs.g = g
		}

		numRows--
		beginAtRow++
	}

	ch <- numGuardsStuck
}

func Part2(fileLines []string) {

	var wg sync.WaitGroup
	ch := make(chan int)

	nrows := len(fileLines)
	const INCR = 1

	for row := 0; row < nrows; row = row + INCR {
		wg.Add(1)

		count := nrows - row

		if count > INCR {
			count = INCR
		}

		go testAddingObstruction(&wg, ch, row, count, fileLines)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	count := 0

	for c := range ch {
		count = count + c
	}

	fmt.Println("Part2: ", count)
}

func main() {
	file, err := os.Open(FILE_TEST)

	if err != nil {
		panic(err)
	}

	defer file.Close()

	lines := getFileLines(file)
	// Part1(lines)
	Part2(lines)
}
