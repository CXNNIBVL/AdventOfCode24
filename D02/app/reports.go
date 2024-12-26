package app

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

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

func parseLinesToReports(lines []string) ([]Report, error) {
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

func isReportSafe(report Report) bool {

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

func EvaluateReports(lines []string) []bool {
	reports, err := parseLinesToReports(lines)

	if err != nil {
		fmt.Fprintf(os.Stderr, "error occurred while parsing the datafile: %s", err.Error())
		os.Exit(1)
	}

	if len(reports) == 0 {
		fmt.Fprintf(os.Stderr, "no reports parsed")
		os.Exit(1)
	}

	safeList := make([]bool, 0, 10)

	for _, report := range reports {
		safeList = append(safeList, isReportSafe(report))
	}

	return safeList
}
