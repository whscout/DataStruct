package main

import (
	"DataStruct/triangle"
	"fmt"
)

func main() {
	var path [4][4]int
	path[0][0] = 2
	path[1][0] = 3
	path[1][1] = 4
	path[2][0] = 6
	path[2][1] = 5
	path[2][2] = 7
	path[3][0] = 4
	path[3][1] = 1
	path[3][2] = 8
	path[3][3] = 3

	triangle.FindMinPath(path, 0, 0)

	fmt.Println(triangle.PathSum)

	fmt.Println(triangle.FindMinPathByDP(path))
}
