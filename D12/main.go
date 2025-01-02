package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
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

type Mat2 [][]rune

func (mat *Mat2) IsInsideBounds(v Vec2) bool {
	return v.x >= 0 && v.x < len((*mat)[0]) && v.y >= 0 && v.y < len((*mat))
}

func (mat *Mat2) Value(v Vec2) (val rune, ok bool) {
	if mat.IsInsideBounds(v) {
		return (*mat)[v.y][v.x], true
	}

	return '\xFF', false
}

func parseMatrix(f *os.File) Mat2 {
	sc := bufio.NewScanner(f)

	mat := [][]rune{}

	for sc.Scan() {
		line := []rune{}
		for _, c := range sc.Text() {
			line = append(line, c)
		}

		mat = append(mat, line)
	}

	return mat
}

type Direction int

const (
	DIR_UP Direction = iota
	DIR_DOWN
	DIR_LEFT
	DIR_RIGHT
	DIR_XX
)

func (d Direction) Inverse() (i Direction) {
	switch d {
	case DIR_UP:
		i = DIR_DOWN
		break
	case DIR_DOWN:
		i = DIR_UP
		break
	case DIR_LEFT:
		i = DIR_RIGHT
		break
	case DIR_RIGHT:
		i = DIR_LEFT
		break
	}

	return
}

func walkRegion(thisId rune, position Vec2, mat Mat2, commandedDirection Direction, visited *[]Vec2) RegionInfo {

	// Mark as visited
	*visited = append(*visited, position)

	thisInfo := RegionInfo{area: 1, perimeter: 3}

	if commandedDirection == DIR_XX {
		thisInfo.perimeter = 4
	}

	evalNext := func(nextPosition Vec2, nextDir Direction) {
		if nextId, inBounds := mat.Value(nextPosition); inBounds && commandedDirection != nextDir.Inverse() && nextId == thisId {
			// Remove border in direction
			thisInfo.perimeter--

			if slices.Contains(*visited, nextPosition) {
				return
			}

			nextInfo := walkRegion(nextId, nextPosition, mat, nextDir, visited)
			thisInfo.area = thisInfo.area + nextInfo.area
			thisInfo.perimeter = thisInfo.perimeter + nextInfo.perimeter
		}
	}

	// Go right
	evalNext(Vec2{x: position.x + 1, y: position.y}, DIR_RIGHT)

	// Go left
	evalNext(Vec2{x: position.x - 1, y: position.y}, DIR_LEFT)

	// Go up
	evalNext(Vec2{x: position.x, y: position.y - 1}, DIR_UP)

	// Go down
	evalNext(Vec2{x: position.x, y: position.y + 1}, DIR_DOWN)

	return thisInfo
}

type RegionInfo struct {
	area, perimeter int
}

func (r *RegionInfo) calculatePrice() int {
	return r.area * r.perimeter
}

func main() {
	f, err := os.Open(TEST_FILE2)

	if err != nil {
		panic(err)
	}

	defer f.Close()

	mat := parseMatrix(f)

	regionInfo := make(map[rune][]RegionInfo)

	var visited []Vec2

	for row, line := range mat {
		for col, r := range line {
			if !slices.Contains(visited, Vec2{x: col, y: row}) {

				info := walkRegion(r, Vec2{x: col, y: row}, mat, DIR_XX, &visited)

				var v []RegionInfo

				if _, exists := regionInfo[r]; exists {
					v = regionInfo[r]
				}

				regionInfo[r] = append(v, info)
			}
		}
	}

	price := 0

	for _, infos := range regionInfo {
		for _, info := range infos {
			price = price + info.calculatePrice()
		}
	}

	fmt.Println(price)
}
