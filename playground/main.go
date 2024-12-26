package main

import (
	"fmt"

	"github.com/CXNNIBVL/goutil/iter"
	"github.com/CXNNIBVL/goutil/math"
)

func main() {

	// Case square nxn
	n := 5

	fmt.Printf("# Case: Square %dx%d Matrix:\n", n, n)

	for x_ := range iter.Interval(-(n - 1), n) {
		meta := struct {
			y, x  int
			items int
		}{
			y: max(-x_, 0), x: max(x_, 0),
			items: n - math.Abs(x_),
		}

		fmt.Printf("%+v\n", meta)
	}

	// Case rectangular
	r, c := 3, 5

	fmt.Printf("# Case: Rectangular %dx%d Matrix:\n", r, c)

	for x_ := range iter.Interval(-(r - 1), c) {
		meta := struct {
			y, x  int
			items int
		}{
			y: max(-x_, 0), x: max(x_, 0),
			items: max(r, c) - math.Abs(max(r-x_, x_)) + min(-x_, r),
		}
		// fmt.Printf("%d - |%d| + %d\n", max(r, c), max(r-x_, x_), min(-x_, 1))
		fmt.Printf("x'= %d, %+v\n", x_, meta)
	}
}

/*
   5x5 mat

   -4 - 1 => x' -> (y, x): x'=-4 -> (4, 0)
   -3 - 2 => x'=-3 -> (3, 0)
   -2 - 3 => x'=-2 -> (2, 0)
   -1 - 4 => x'=-1 -> (1, 0)
    0 - 5 => x'=0 -> (0,0)
    1 - 4 => x'=1 -> (0, 1)
    2 - 3 => x'=2 -> (0, 2)
    3 - 2 => x'=3 -> (0,3)
    3 - 1 == x'=4 -> (0,4)

    => Flip x,y

   3x5 mat
*/
