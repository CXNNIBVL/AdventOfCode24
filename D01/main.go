package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

const (
	INPUT_FILE = "../inputs/D01/input"
)

func sortList(list []uint) {
	sort.Slice(list, func(i, j int) bool {
		return list[i] < list[j]
	})
}

func getTotalDistance(leftList, rightList []uint) uint {

	var totalDistance uint = 0

	for ix := 0; ix < len(leftList); ix++ {
		var diff, r, l uint = 0, rightList[ix], leftList[ix]

		if r < l {
			diff = l - r
		} else {
			diff = r - l
		}

		totalDistance = totalDistance + diff
	}

	return totalDistance
}

func parseFile(f *os.File) (left, right []uint) {
	sc := bufio.NewScanner(f)

	left, right = make([]uint, 0, 100), make([]uint, 0, 100)

	for sc.Scan() {
		leftIntStr, spaceAndIntStr, _ := strings.Cut(sc.Text(), " ")

		l, _ := strconv.ParseUint(leftIntStr, 10, 64)

		rightIntStr := strings.ReplaceAll(spaceAndIntStr, " ", "")

		r, _ := strconv.ParseUint(rightIntStr, 10, 64)

		left = append(left, uint(l))
		right = append(right, uint(r))
	}

	return left, right
}

func Part1(left, right []uint) uint {
	sortList(left)
	sortList(right)

	return getTotalDistance(left, right)
}

func Part2(left, right []uint) uint64 {
	multMap := make(map[uint]uint)

	for _, v := range left {
		multMap[v] = 0
	}

	for _, v := range right {
		m, ok := multMap[v]

		if !ok {
			continue
		}

		multMap[v] = m + 1
	}

	var similarityScore uint64 = 0

	for k, v := range multMap {
		similarityScore = similarityScore + uint64(k)*uint64(v)
	}

	return similarityScore
}

func parseInputs() (left, right []uint) {
	input, err := os.Open(INPUT_FILE)

	if err != nil {
		panic(err)
	}

	defer input.Close()

	left, right = parseFile(input)
	return left, right
}

func main() {

	left, right := parseInputs()
	sol1 := Part1(left, right)
	similarityScore := Part2(left, right)

	fmt.Printf("Part1: Total Distance = %d\n", sol1)
	fmt.Printf("Part2: Similarity Score = %d\n", similarityScore)
}
