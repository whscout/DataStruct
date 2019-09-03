package LongestIncresingSeqeuce

import (
	"fmt"
	"math"
)

/*
最长上升子序列
动态规划

状态： LIS[i] 表示 0 - i（包括元素i）中最长的上升子序列
状态转移方程 LIS[i] = max(LIS[0]...LIS[j]) + 1 (j 0 -> i - 1 && arr[i] > arr[j])
 */

func LIS(arr []int) int {
	max := 0

	if len(arr) <= 1 {
		return 1
	}
	dp := make([]int, len(arr))
	for i := 0; i < len(arr); i++ {
		dp[i] = 1
	}

	for i := 1; i < len(arr); i++ {
		for j := 0; j < i; j++ {
			if arr[i] > arr[j] {
				dp[i] = int(math.Max(float64(dp[i]), float64(dp[j] + 1)))
			}
		}
		max = int(math.Max(float64(max), float64(dp[i])))
	}
	fmt.Println(dp)

	return max
}

/*
二分查找
维护一个最大上升子序列数组 lis

首先遍历整个数组
在维护的最大上升子序列中查找
若比最大的元素都大 则append到lis数组中
否则找到第一个比他大的元素替换掉

 */
func BoundLIS(arr []int) int {

	lis := make([]int, 1)
	lis[0] = arr[0]

	for i := 1; i < len(arr); i++ {
		index := lower_bound(lis, 0, len(lis), arr[i])
		fmt.Println("index", index)
		if len(lis) == index {
			// 添加
			lis = append(lis, arr[i])
		} else {
			// 替换
			lis[index] = arr[i]
		}
	}
	fmt.Println(lis)
	return len(lis)
}

func lower_bound(arr []int, start, end int, num int) int {
	index := -1


	for start < end {
		mid := (start + end) / 2
		fmt.Println(arr, start, end, mid)

		if num == arr[mid] {
			break
		}
		if num > arr[mid] {
			start = mid + 1
		} else if num < arr[mid] {
			end = mid - 1
		}
	}

	index = (start + end) / 2
	if index < 0 {
		index = 0
	}

	return index
}
