package main

import (
	"testing"
)

func BenchmarkPart1(b *testing.B) {
	for n := 0; n < b.N; n++ {
		b.StopTimer()
		left, right := parseInputs()
		b.StartTimer()
		Part1(left, right)
	}
}

func BenchmarkPart2(b *testing.B) {
	for n := 0; n < b.N; n++ {
		b.StopTimer()
		left, right := parseInputs()
		Part1(left, right)
		b.StartTimer()
		Part2(left, right)
	}
}
