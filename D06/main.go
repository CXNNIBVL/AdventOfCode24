package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
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
	m     Mat2
	g     Guard
	hm    map[Vec2]int
	hasHm bool
}

func (gs *GameState) loop(tickCb func(gs *GameState) bool) {
	for {
		// Heatmap
		if gs.hasHm {
			gs.hm[gs.g.position]++
		}

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

func Part1(fileLines []string) []Vec2 {
	m, guard := parseMapFromLines(fileLines)
	gs := GameState{m: m, g: guard}

	gs.addHeatMap()
	gs.loop(func(gs *GameState) bool {
		return false
	})

	fmt.Println("Part1: ", len(gs.hm))

	uniqPoints := []Vec2{}
	for k := range gs.hm {
		uniqPoints = append(uniqPoints, k)
	}

	return uniqPoints
}

func (gs *GameState) addHeatMap() {
	gs.hasHm = true
	gs.hm = make(map[Vec2]int)
}

func testAddingObstructionV2(uniqPoints []Vec2, fileLines []string, wg *sync.WaitGroup, ch chan<- int) {
	defer wg.Done()

	newFileLines := []string{}
	for _, line := range fileLines {
		newFileLines = append(newFileLines, strings.Clone(line))
	}

	m, g := parseMapFromLines(newFileLines)
	gs := GameState{m: m, g: g}

	numGuardsStuck := 0

	for _, p := range uniqPoints {

		obj := inspectMapIndex(p.y, p.x, gs.m)

		if obj == OBJ_OBSTACLE || (p.y == gs.g.position.y && p.x == gs.g.position.x) {
			// skip
			continue
		}

		// Place obstacle
		gs.m[p.y][p.x] = OBJ_OBSTACLE

		stuck := false
		steps := 0
		gs.loop(func(gs *GameState) bool {
			steps = steps + 1

			if steps > 25000 {
				stuck = true
			}

			return stuck
		})

		if stuck {
			numGuardsStuck = numGuardsStuck + 1
		}

		// Cleanup placed obstacle
		gs.m[p.y][p.x] = obj
		// Reset guard
		gs.g = g
	}

	ch <- numGuardsStuck
}

func Part2V2(fileLines []string, uniqPoints []Vec2) {

	var wg sync.WaitGroup
	ch := make(chan int)

	const INCR = 30

	id := 0
	tmp := uniqPoints
	for len(tmp) > 0 {
		wg.Add(1)

		count := INCR

		if len(tmp) < INCR {
			count = len(tmp)
		}

		go testAddingObstructionV2(tmp[0:count], fileLines, &wg, ch)
		tmp = tmp[count:]
		id++
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
	file, err := os.Open(FILE)

	if err != nil {
		panic(err)
	}

	defer file.Close()

	lines := getFileLines(file)
	uniqPoints := Part1(lines)
	Part2V2(lines, uniqPoints)
}
