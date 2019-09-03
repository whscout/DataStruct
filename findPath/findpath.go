package findPath

import "fmt"

var Row int
var Col int
var Mem [8][8]int


func Find(gird [8][8]int, row, col int) int {
	if !isValid(gird, row, col) {
		return 0
	}
	if isEnd(row, col) {
		return 1
	}
	if Mem[row][col] == 0 {
		Mem[row][col] = Find(gird, row + 1, col) + Find(gird, row, col + 1)
	}
	//fmt.Println(Mem[row][col])
	return Mem[row][col]
}

func isValid(gird [8][8]int, row, col int) bool {
	//fmt.Println(row, col)
	if row >= Row || col >= Col {
		return false
	}
	if gird[row][col] == 0 {
		return true
	} else {
		return false
	}
}

func isEnd(row, col int) bool {
	if row == (Row - 1) && col == (Col - 1) {
		return true
	} else {
		return false
	}
}

func FindDY(gird [8][8]int) int {
	for row := 6; row >= 0; row-- {
		for col := 6; col >= 0; col-- {
			fmt.Println(row, col, gird[row][col])
			if gird[row][col] == 0 {
				//fmt.Println(Mem[row -1][col], Mem[row][col - 1])
				Mem[row][col] = Mem[row + 1][col] + Mem[row][col + 1]
			}
		}
	}
	fmt.Println(Mem)
	return Mem[0][0]
}


