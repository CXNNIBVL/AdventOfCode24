package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	FILE      = "../inputs/D07/input"
	TEST_FILE = "../inputs/D07/input_test"
)

func readAllLines(f *os.File) []string {
	lines := []string{}

	sc := bufio.NewScanner(f)

	for sc.Scan() {
		lines = append(lines, sc.Text())
	}

	return lines
}

type Equation struct {
	result  int64
	numbers []int64
}

func parseEqsFromLines(lines []string) []Equation {
	eqs := []Equation{}

	for _, line := range lines {
		s := strings.Split(line, ":")

		resultEqStr := s[0]

		after, _ := strings.CutPrefix(s[1], " ")

		nums := strings.Split(after, " ")

		result, _ := strconv.ParseInt(resultEqStr, 10, 64)

		eq := Equation{
			result:  result,
			numbers: []int64{},
		}

		for _, num := range nums {
			n, _ := strconv.ParseInt(num, 10, 64)
			eq.numbers = append(eq.numbers, n)
		}

		eqs = append(eqs, eq)
	}

	return eqs
}

func recurse(res int64, ix int, nums []int64, currentVal int64, addConcat bool) bool {
	if ix >= len(nums) {
		return currentVal == res
	}

	// Test addition
	if recurse(res, ix+1, nums, currentVal+nums[ix], addConcat) {
		return true
	}

	// Test multiplication
	if recurse(res, ix+1, nums, currentVal*nums[ix], addConcat) {
		return true
	}

	// Test concatenation
	if addConcat {
		a, b := strconv.FormatInt(currentVal, 10), strconv.FormatInt(nums[ix], 10)
		v, _ := strconv.ParseInt(a+b, 10, 64)
		if recurse(res, ix+1, nums, v, addConcat) {
			return true
		}
	}

	return false
}

func isValidEq(res int64, nums []int64, addConcat bool) bool {
	return recurse(res, 1, nums, nums[0], addConcat)
}

func Part1(eqs []Equation) int64 {
	var sum int64 = 0
	for _, eq := range eqs {
		if isValidEq(eq.result, eq.numbers, false) {
			sum = sum + eq.result
		}
	}

	return sum
}

func Part2(eqs []Equation) int64 {
	var sum int64 = 0
	for _, eq := range eqs {
		if isValidEq(eq.result, eq.numbers, true) {
			sum = sum + eq.result
		}
	}

	return sum
}

func main() {
	file, err := os.Open(FILE)

	if err != nil {
		panic(err)
	}

	defer file.Close()

	lines := readAllLines(file)
	eqs := parseEqsFromLines(lines)

	sum := Part1(eqs)
	fmt.Println("Part1: ", sum)
	sum = Part2(eqs)
	fmt.Println("Part2: ", sum)
}
