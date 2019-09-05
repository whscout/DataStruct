package sqrt

import (
"fmt"
"math"
)

func Binary(n float64) float64 {
	max := n
	min := n / 2
	error := 1e-10
	i := 0
	for math.Abs(max - min) > error {
		if min * min > n {
			max = min
			min = min / 2
		}
		if min * min < n {
			min = (max + min) /  2
		}
		i++
		//fmt.Println(min, max)
	}
	fmt.Println(i)
	return max
}

func Newton(n float64) float64 {
	x0 := n / 2
	x1 := (x0 + (n / x0)) / 2
	error := 1e-10
	i := 0
	for math.Abs(x0 - x1) > error {
		x0 = x1
		x1 = (x0 + (n / x0)) / 2
		i++
	}

	fmt.Println(i)
	return x0
}

