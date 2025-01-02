package main

import (
	"testing"
)

func BenchmarkPart1(b *testing.B) {
	for n := 0; n < b.N; n++ {
		b.StopTimer()
		mat := parseInputs()

		b.StartTimer()
		Part1(mat)
	}
}

func BenchmarkPart2(b *testing.B) {
	for n := 0; n < b.N; n++ {
		b.StopTimer()
		mat := parseInputs()

		Part1(mat)
		b.StartTimer()
		Part2(mat)
	}
}
