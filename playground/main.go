package main

import (
	"fmt"

	"github.com/CXNNIBVL/goutil/math"
)

type Vec2 struct {
	x, y int
}

func main() {
	v1 := Vec2{x: 1, y: 2}
	v2 := Vec2{x: -3, y: -2}

	slope := (v2.y - v1.y) / (v2.x - v1.x)
	if slope == 1 {
		fmt.Println("Valid")
	}

	sameHorizAndVertDistance := math.Abs(v2.y-v1.y) == math.Abs(v2.x-v1.x)

	if sameHorizAndVertDistance {
		fmt.Println("Valid")
	}
}
