package main

import "fmt"

func pow(x float64, n int) float64 {
	if n < 0 {
		x = 1 / x
		n = -n
	}
	pow := 1.0

	for n > 0 {
		if n & 1 == 1 {
			pow *= x
		}
		x *= x
		n >>= 1
	}

	return pow
}

func main() {
	fmt.Print(pow(5, -2))
}
