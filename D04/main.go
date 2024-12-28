package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/CXNNIBVL/goutil/iter"
	"github.com/CXNNIBVL/goutil/math"
)

const (
	FILE      = "input"
	TEST_FILE = "input_test3"
	SEQ       = "XMAS"
)

type Mat2 []string

func parseFileAsMatrix(f *os.File) Mat2 {
	mat := make(Mat2, 0, 10)
	sc := bufio.NewScanner(f)

	for sc.Scan() {
		mat = append(mat, sc.Text())
	}

	return mat
}

func (m *Mat2) NumRowsCols() (rows, cols int) {
	return len(*m), len((*m)[0])
}

func (m *Mat2) Reverse() {
	for ix := range *m {
		(*m)[ix] = reverseString((*m)[ix])
	}
}

func reverseString(s string) (reversed string) {
	for _, c := range s {
		reversed = string(c) + reversed
	}
	return
}

func countSequenceMatches(str, seq string, addReverseSearch bool) int {
	count := strings.Count(str, seq)

	if addReverseSearch {
		count = count + strings.Count(reverseString(str), seq)
	}

	return count
}

func countHorizontalMatches(mat Mat2, seq string, addScanReverse bool) int {
	matches := 0

	for _, row := range mat {
		matches = matches + countSequenceMatches(row, seq, addScanReverse)
	}

	return matches
}

func countVerticalMatches(mat Mat2, seq string, addScanReverse bool) int {
	matches := 0

	_, ncols := mat.NumRowsCols()

	for x := range iter.Interval(0, ncols) {
		var column string
		for _, row := range mat {
			column = column + string(row[x])
		}
		matches = matches + countSequenceMatches(column, seq, addScanReverse)
	}

	return matches
}

func countDiagonalMatches(mat Mat2, seq string, addScanReverse bool) int {
	matches := 0

	_, ncols := mat.NumRowsCols()

	for x_ := range iter.Interval(-(ncols - 1), ncols) {
		y, x, items := max(-x_, 0), max(x_, 0), ncols-math.Abs(x_)

		var diag string
		for i := range iter.Interval(0, items) {
			diag = diag + string(mat[y+i][x+i])
		}

		matches = matches + countSequenceMatches(diag, seq, addScanReverse)
	}

	return matches
}

func Part1(mat Mat2) {
	if rows, cols := mat.NumRowsCols(); rows != cols {
		panic("Cannot handle non quadratic search matrices")
	}

	seq := SEQ

	matches := countHorizontalMatches(mat, seq, true)
	matches = matches + countVerticalMatches(mat, seq, true)
	matches = matches + countDiagonalMatches(mat, seq, true)

	mat.Reverse()

	matches = matches + countDiagonalMatches(mat, seq, true)

	mat.Reverse()

	fmt.Printf("Part1: Found %d matches\n", matches)
}

func main() {

	file, err := os.Open(FILE)

	if err != nil {
		fmt.Println(err)
		return
	}

	defer file.Close()

	mat := parseFileAsMatrix(file)

	Part1(mat)
}
