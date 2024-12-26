package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

const (
	DATA_FILE = "input"
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

type Report []int

func parseLineToReport(line string) (Report, error) {

	items := strings.Split(line, " ")

	if len(items) == 0 {
		return nil, fmt.Errorf("report was empty")
	}

	report := make(Report, 0, 10)

	for _, item := range items {
		parsed, err := strconv.ParseInt(item, 10, 64)

		if err != nil {
			return nil, err
		}

		report = append(report, int(parsed))
	}

	return report, nil
}

func ParseLinesToReports(lines []string) ([]Report, error) {
	if len(lines) == 0 {
		return nil, fmt.Errorf("no lines were read")
	}

	reports := make([]Report, 0, 5)

	for ix, line := range lines {
		report, err := parseLineToReport(line)

		if err != nil {
			fmt.Fprintf(os.Stderr, "error occurred when parsing datafile line %d: %s", ix, err.Error())
			os.Exit(1)
		}

		if len(report) == 0 {
			fmt.Fprintf(os.Stderr, "report for line %d is empty", ix)
			os.Exit(1)
		}

		reports = append(reports, report)
	}

	return reports, nil
}

func IsReportSafe(report Report) bool {

	allIncreasing, allDecreasing, adjacentDiffOk := true, true, true

	for ix := range report[0 : len(report)-1] {
		current, next := report[ix], report[ix+1]
		if allIncreasing && current > next {
			allIncreasing = false
		}

		if allDecreasing && current < next {
			allDecreasing = false
		}

		if !allIncreasing && !allDecreasing {
			// Abort
			return false
		}

		diff := math.Abs(float64(current - next))
		if adjacentDiffOk && !(diff >= 1 && diff <= 3) {
			// Abort
			return false
		}
	}

	return (allIncreasing || allDecreasing) && adjacentDiffOk
}

func Part1(reports []Report) {
	safeCtr := 0

	for _, r := range reports {
		if IsReportSafe(r) {
			safeCtr = safeCtr + 1
		}
	}

	fmt.Printf("Part 1: Safe reports = %d\n", safeCtr)
}

func tryApplyProblemDampener(report Report) bool {

	for ix := range report {
		reportCopy := slices.Clone(report)
		reportCopy = slices.Delete(reportCopy, ix, ix+1)
		if IsReportSafe(reportCopy) {
			return true
		}
	}

	return false
}

func Part2(reports []Report) {
	safeCtr := 0

	for _, report := range reports {

		if IsReportSafe(report) {
			safeCtr = safeCtr + 1
			continue
		}

		ok := tryApplyProblemDampener(report)

		if ok {
			safeCtr = safeCtr + 1
		}
	}

	fmt.Printf("Part 2: Safe reports = %d\n", safeCtr)
}

func main() {

	lines, err := readDataFileLines()

	if err != nil {
		panic(err)
	}

	reports, err := ParseLinesToReports(lines)

	if err != nil {
		panic(err)
	}

	Part1(reports)
	Part2(reports)
}
