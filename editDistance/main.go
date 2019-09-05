package main

import (
	"fmt"
	"math"
)

// word1 => word2 最少需要的步骤(insert delete replace三种操作)
// 状态： DP[i, j] 表示word1的前i个字符 变为Word2的前j个字符需要的最少步骤

// 状态方程
// DP[i, j]
// if word1[i] == word2[j] : DP[i, j] = DP[i - 1, j - 1]
// else :
// DP[i, j] = min(DP[i, j - 1], DP[i - 1, j], DP[i - 1, j - 1])

// 关于初始化DP  DP[i][0] 表示Word1 的第i个字符 匹配到word2的第0个字符 至少需要i步
// 同理 DP[0][j] 表示word1 的第0个字符 匹配到word2的第j个字符 至少需要j步

func distance(word1, word2 []rune) int {
	n, m := len(word1), len(word2)
	var dp [][]int
	for i := 0; i < n; i++ {
		tmp := make([]int, m)
		dp = append(dp, tmp)
	}

	// 初始化数据
	for i := 0; i < n; i++ {
		dp[i][0] = i
	}

	for j := 0; j < m; j++ {
		dp[0][j] = j
	}

	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			if word1[i] == word2[j] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = int(math.Min(math.Min(float64(dp[i-1][j]), float64(dp[i][j-1])), float64(dp[i-1][j-1]))) + 1
			}
		}
	}

	return dp[n - 1][m - 1]
}

func main() {
	word1 := []rune("intention")
	word2 := []rune("nation")

	fmt.Println(distance(word1, word2))
}
