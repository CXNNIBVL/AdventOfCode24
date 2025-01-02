package main

import "testing"

func BenchmarkPart1And2(b *testing.B) {
	for n := 0; n < b.N; n++ {
		b.StopTimer()
		updates, rules := parseInputs()
		b.StartTimer()
		Part1And2(updates, rules)
	}
}
