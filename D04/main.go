package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	FILE      = "input"
	TEST_FILE = "input_test3"
	SEQ       = "XMAS"
)

func parseFileAsMatrix(f *os.File) []string {
	mat := make([]string, 0, 10)
	sc := bufio.NewScanner(f)

	for sc.Scan() {
		mat = append(mat, sc.Text())
	}

	return mat
}

func reduceMatToQuadMat(startX, startY int, mat []string, width int) []string {

	nrows, ncols := len(mat), len(mat[0])

	yBound := startY + width
	xBound := startX + width

	// Submatrix can't fit
	if yBound > nrows || xBound > ncols {
		return nil
	}

	newMat := make([]string, width)

	newY := 0
	oldY := startY
	for newY < width {
		newMat[newY] = strings.Clone(mat[oldY])
		newY++
		oldY++
	}

	// For each row, truncate columns
	for ix := range newMat {
		newMat[ix] = newMat[ix][startX:xBound]
	}

	return newMat
}

func getTopLeftToBottomRightDiagonalOfQuadMat(mat []string) string {

	width := len(mat)

	var s string

	for ix := 0; ix < width; ix++ {
		s = s + string(mat[ix][ix])
	}

	return s
}

func getBottomLeftToTopRightDiagonalOfQuadMat(mat []string) string {
	width := len(mat)

	var s string

	y := width - 1
	x := 0

	for x < width {
		s = s + string(mat[y][x])
		y = y - 1
		x = x + 1
	}

	return s
}

func countSequenceMatches(str, seq string, addReverseSearch bool) int {
	count := strings.Count(str, seq)

	if addReverseSearch {
		var reversed string

		for _, c := range str {
			reversed = string(c) + reversed
		}

		count = count + strings.Count(reversed, seq)
	}

	return count
}

func divideMatIntoQuadMatsAndCountMatches(mat []string, seq string, searchDiag, searchHoriz, searchVert bool) int {

	width := len(seq)
	nrows := len(mat)
	ncols := len(mat[0])

	matches := 0

	for y := 0; y < nrows; y++ {
		for x := 0; x < ncols; x++ {
			submat := reduceMatToQuadMat(x, y, mat, width)

			if submat == nil {
				continue
			}

			if len(submat) != len(seq) || len(submat[0]) != len(seq) {
				panic("Non square matrix detected")
			}

			if searchDiag {
				tlbr := getTopLeftToBottomRightDiagonalOfQuadMat(submat)
				bltr := getBottomLeftToTopRightDiagonalOfQuadMat(submat)

				// Diagonal
				matches = matches + countSequenceMatches(tlbr, seq, true)
				matches = matches + countSequenceMatches(bltr, seq, true)
			}

			if searchHoriz {
				for _, row := range submat {
					matches = matches + countSequenceMatches(row, seq, true)
				}
			}

			if searchVert {
				for subx := 0; subx < width; subx++ {
					var colstr string
					for suby := 0; suby < width; suby++ {
						colstr = colstr + string(submat[suby][subx])
					}
					matches = matches + countSequenceMatches(colstr, seq, true)
				}
			}
		}
	}

	matFitsAllBlocks := nrows%width == 0 && ncols%width == 0

	if !matFitsAllBlocks {
		panic("Can only handle search matrices where row size and col size are a multiple of the queries length")
	}

	return matches
}

func main() {

	file, err := os.Open(TEST_FILE)
	seq := SEQ

	if err != nil {
		fmt.Println(err)
		return
	}

	defer file.Close()

	mat := parseFileAsMatrix(file)

	matches := divideMatIntoQuadMatsAndCountMatches(mat, seq, true, true, true)

	fmt.Println(matches)
}
