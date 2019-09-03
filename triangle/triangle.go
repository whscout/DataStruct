package triangle

import (
	"fmt"
	"math"
)

var PathSum []int
var sum int = 0
func FindMinPath(path [4][4]int, i, j int) {
	sum += path[i][j]
	if i == 3 {
		PathSum = append(PathSum, sum)
		return
	}
	FindMinPath(path, i + 1, j)
	sum -= path[i + 1][j]

	FindMinPath(path, i + 1, j + 1)
	sum -= path[i + 1][j + 1]
}

func FindMinPathByDP(path [4][4]int) int {
	var opt [4][4]int
	opt[3][0] = path[3][0]
	opt[3][1] = path[3][1]
	opt[3][2] = path[3][2]
	opt[3][3] = path[3][3]

	for i := len(path) - 2; i >= 0; i-- {
		for j := 0; j < len(path[i]) - 1; j++ {
			fmt.Println(i, j)
			opt[i][j] = path[i][j] + int(math.Min(float64(opt[i + 1][j]), float64(opt[i + 1][j + 1])))
		}
	}

	return opt[0][0]
}
