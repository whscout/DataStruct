package main

import (
	"DataStruct/LongestIncresingSeqeuce"
	"fmt"
)

func main() {
	arr := []int{10, 9, 2, 5, 3, 7, 101, 18, 20}

	fmt.Println(LongestIncresingSeqeuce.LIS(arr))

	//fmt.Println(LongestIncresingSeqeuce.BoundLIS(arr))
}
