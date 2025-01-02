package main

import "testing"

func BenchmarkPart1(b *testing.B) {
	for n := 0; n < b.N; n++ {
		b.StopTimer()
		lines, err := readDataFileLines()

		if err != nil {
			panic(err)
		}

		reports, err := ParseLinesToReports(lines)

		if err != nil {
			panic(err)
		}

		b.StartTimer()
		Part1(reports)
	}
}

func BenchmarkPart2(b *testing.B) {
	for n := 0; n < b.N; n++ {
		b.StopTimer()
		lines, err := readDataFileLines()

		if err != nil {
			panic(err)
		}

		reports, err := ParseLinesToReports(lines)

		if err != nil {
			panic(err)
		}

		Part1(reports)
		b.StartTimer()
		Part2(reports)
	}
}
