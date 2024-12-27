package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func parseFileAsMatrix(f *os.File) []string {
	mat := make([]string, 0, 10)
	sc := bufio.NewScanner(f)

	for sc.Scan() {
		mat = append(mat, sc.Text())
	}

	return mat
}

func reduceToQuadMatrix(startX, startY int, mat []string, width int, matRowSize int, matColSize int) []string {

	ybound := startY + width
	xbound := startX + width

	// Submatrix can fit
	if ybound > matRowSize || xbound > matColSize {
		return nil
	}

	newMat := make([]string, width)

	ny := 0
	oy := startY
	for ny < width {
		newMat[ny] = strings.Clone(mat[oy])
		ny++
		oy++
	}

	// For each row, truncate columns
	for ix := range newMat {
		newMat[ix] = newMat[ix][startX:xbound]
	}

	return newMat
}

func getTopLeftToBottomRightDiagonalOfQuadMat(mat []string) string {

	width := len(mat)

	var sb strings.Builder

	for ix := 0; ix < width; ix++ {
		sb.WriteByte(mat[ix][ix])
	}

	return sb.String()
}

func getBottomLeftToTopRightDiagonalOfQuadMat(mat []string) string {
	width := len(mat)

	var sb strings.Builder

	y := width - 1
	x := 0

	for x < width {
		sb.WriteByte(mat[y][x])
		y = y - 1
		x = x + 1
	}

	return sb.String()
}

func countMatchingSequences(str, seq string, addReverseSearch bool) int {
	count := 0

	count = count + strings.Count(str, seq)

	if addReverseSearch {
		var sb strings.Builder

		for i := len(str) - 1; i >= 0; i-- {
			sb.WriteByte(str[i])
		}

		count = count + strings.Count(sb.String(), seq)
	}

	return count
}

func divideIntoQuadMatricesAndCountMatches(mat []string, seq string) int {

	width := len(seq)

	matFitsAllBlocks := len(mat)%width == 0 && len(mat[0])%width == 0

	if !matFitsAllBlocks {
		panic("Still TODO")
		// TODO:
	}

	matches := 0

	nrows := len(mat)
	ncol := len(mat[0])

	for y := 0; y < nrows; y++ {
		for x := 0; x < ncol; x++ {
			submat := reduceToQuadMatrix(x, y, mat, width, nrows, ncol)

			if submat == nil {
				continue
			}

			if len(submat) != len(seq) || len(submat[0]) != len(seq) {
				panic("Non square matrix detected")
			}

			tlbr := getTopLeftToBottomRightDiagonalOfQuadMat(submat)
			bltr := getBottomLeftToTopRightDiagonalOfQuadMat(submat)

			// Diagonal
			matches = matches + countMatchingSequences(tlbr, seq, true)
			matches = matches + countMatchingSequences(bltr, seq, true)

			// Horizontal
			for _, row := range submat {
				matches = matches + countMatchingSequences(row, seq, true)
			}

			// Vertical
			for sx := 0; sx < width; sx++ {
				var sb strings.Builder
				for sy := 0; sy < width; sy++ {
					sb.WriteByte(submat[sy][sx])
				}
				matches = matches + countMatchingSequences(sb.String(), seq, true)
			}
		}
	}

	return matches
}

func main() {

	// file, err := os.Open(os.Args[1])
	// seq := os.Args[2]

	file, err := os.Open("../input_test2")
	seq := "XMAS"

	if err != nil {
		fmt.Println(err)
		return
	}

	defer file.Close()

	mat := parseFileAsMatrix(file)

	matches := divideIntoQuadMatricesAndCountMatches(mat, seq)

	fmt.Println(matches)
}
