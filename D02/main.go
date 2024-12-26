package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/CXNNIBVL/AdventOfCode24/D02/app"
)

const (
	DATA_FILE = "data.txt"
)

func readDataFileLines() ([]string, error) {
	dataFile, err := os.Open(DATA_FILE)

	if err != nil {
		return nil, err
	}

	defer dataFile.Close()

	dataScanner := bufio.NewScanner(dataFile)

	lines := make([]string, 0, 10)

	for dataScanner.Scan() {
		lines = append(lines, dataScanner.Text())
	}

	return lines, nil
}

func main() {

	lines, err := readDataFileLines()

	if err != nil {
		fmt.Fprintf(os.Stderr, "error occurred while opening datafile: %s", err.Error())
		os.Exit(1)
	}

	safeList := app.EvaluateReports(lines)

	fmt.Printf("Safe Report List: %v", safeList)
}
