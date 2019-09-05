package main

import "fmt"

func findDuplicateNum(data []int) {
	for i := 0; i < len(data); i++ {
		for data[i] != i {
			if data[i] != data[data[i]] {
				data[i], data[data[i]] = data[data[i]], data[i]
			} else {
				fmt.Println(data[i])
				i++
			}
		}
	}
}

func main() {
	data := []int{4, 3, 5, 2, 5, 3}

	findDuplicateNum(data)
}

