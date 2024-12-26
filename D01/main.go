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
	INPUT_FILE = "input"
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

func parseFile(f *os.File) (left, right []uint, e error) {
	sc := bufio.NewScanner(f)

	left, right = make([]uint, 0, 100), make([]uint, 0, 100)

	for sc.Scan() {
		leftIntStr, spaceAndIntStr, found := strings.Cut(sc.Text(), " ")

		if !found {
			return nil, nil, fmt.Errorf("invalid file format")
		}

		l, err := strconv.ParseUint(leftIntStr, 10, 64)

		if err != nil {
			return nil, nil, err
		}

		rightIntStr := strings.ReplaceAll(spaceAndIntStr, " ", "")

		r, err := strconv.ParseUint(rightIntStr, 10, 64)

		if err != nil {
			return nil, nil, err
		}

		left = append(left, uint(l))
		right = append(right, uint(r))
	}

	return left, right, nil
}

func Part1(left, right []uint) {
	fmt.Printf("Part1: Total Distance = %d\n", getTotalDistance(left, right))
}

func advanceSlice[T any](s []T) ([]T, bool) {
	if len(s) == 1 {
		return s, false
	}

	return s[1:], true
}

func Part2(left, right []uint) {
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

	fmt.Printf("Part2: Similarity Score = %d\n", similarityScore)
}

func main() {
	input, err := os.Open(INPUT_FILE)

	if err != nil {
		panic(err)
	}

	defer input.Close()

	left, right, err := parseFile(input)

	if err != nil {
		panic(err)
	}

	sortList(left)
	sortList(right)

	Part1(left, right)
	Part2(left, right)
}
