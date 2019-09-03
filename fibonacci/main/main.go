package main

import (
	"fmt"
	"time"
)

var mem [1000]int

func fibMem(n int) int {
	if n <= 1 {
		return mem[n]
	}
	if mem[n] == 0 {
		mem[n] = fibMem(n - 1) + fibMem(n - 2)
	}
	return mem[n]
}

func fibDY(n int) int {
	for i := 2; i <= n; i++ {
		mem[i] = mem[i - 1] + mem[i - 2]
	}
	return mem[n]
}

func main() {
	mem[0] = 0
	mem[1] = 1
	t1 := time.Now()
	fmt.Println(fibDY(500))
	elapsed := time.Since(t1)

	fmt.Println(elapsed)

	t1 = time.Now()
	fmt.Println(fibMem(50))
	elapsed = time.Since(t1)

	fmt.Println(elapsed)
}
