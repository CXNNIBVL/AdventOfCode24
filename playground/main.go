package main

import (
	"fmt"
	"math"
)

func main() {
	numDigits := 0
	stonenum := 100000
	for s := stonenum; s != 0; s = s / 10 {
		numDigits++
	}

	fmt.Printf("%d %v\n", stonenum, numDigits%2 == 0)
	hi := stonenum / int(math.Pow10(numDigits/2))
	lo := stonenum - hi*int(math.Pow10(numDigits/2))
	fmt.Println("Hi: ", hi)
	fmt.Println("Lo: ", lo)
}
