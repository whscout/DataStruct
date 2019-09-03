package main

import (
	"DataStruct/findPath"
	"fmt"
	"time"
)

func main() {
	findPath.Col = 8
	findPath.Row = 8
	var gird [8][8]int
	gird[1][2] = 1
	gird[1][6] = 1
	gird[2][4] = 1
	gird[3][0] = 1
	gird[3][2] = 1
	gird[3][5] = 1
	gird[4][2] = 1
	gird[5][3] = 1
	gird[5][4] = 1
	gird[5][6] = 1
	gird[6][1] = 1
	gird[6][5] = 1

	findPath.Mem[7][6] = 1
	findPath.Mem[7][5] = 1
	findPath.Mem[7][4] = 1
	findPath.Mem[7][3] = 1
	findPath.Mem[7][2] = 1
	findPath.Mem[7][1] = 1
	findPath.Mem[7][0] = 1

	findPath.Mem[6][7] = 1
	findPath.Mem[5][7] = 1
	findPath.Mem[4][7] = 1
	findPath.Mem[3][7] = 1
	findPath.Mem[2][7] = 1
	findPath.Mem[1][7] = 1
	findPath.Mem[0][7] = 1

	t1 := time.Now()
	fmt.Println(findPath.FindDY(gird))
	fmt.Println(time.Since(t1))
}
