package main

import (
	"DataStruct/NQueens"
	"fmt"
)

func main() {
	NQueens.NQueensDFS(8, 0, 0, 0, 0)

	fmt.Println(NQueens.Count)
}
