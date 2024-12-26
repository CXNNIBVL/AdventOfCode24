package app_test

import (
	"strconv"
	"testing"

	"github.com/CXNNIBVL/AdventOfCode24/D02/app"
)

func Test_AdventOfCode(t *testing.T) {

	lines := []string{
		"7 6 4 2 1",
		"1 2 7 8 9",
		"9 7 6 2 1",
		"1 3 2 4 5",
		"8 6 4 4 1",
		"1 3 6 7 9",
	}

	reports := app.EvaluateReports(lines)

	expect := []bool{
		true,
		false,
		false,
		false,
		false,
		true,
	}

	if len(reports) != len(expect) {
		t.Fatal("Number of reports doesn't equal the expected length")
	}

	for ix := range expect {
		e, r := strconv.FormatBool(expect[ix]), strconv.FormatBool(reports[ix])
		if e != r {
			t.Fatalf("Report %d doesn't equal the expected report. Expected \"%s\", found \"%s\")", ix, e, r)
		}
	}
}
