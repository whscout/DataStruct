package main

import (
	"fmt"
	"sync"
)

func quickSort(data []int, left, right int) {

	wg := sync.WaitGroup{}
	tmp := data[left]
	p := left
	i, j := left, right

	for i <= j {
		for j >= p && data[j] >= tmp {
			j--
		}
		if j >= p {
			data[p] = data[j]
			p = j
		}
		for i <= p && data[i] <= tmp {
			i++
		}
		if i <= p {
			data[p] = data[i]
			p = i
		}
	}
	data[p] = tmp

	if p > left {
		wg.Add(1)
		go func() {
			quickSort(data, left, p - 1)
			wg.Done()
		}()
	}
	if right > p {
		wg.Add(1)
		go func() {
			quickSort(data, p + 1, right)
			wg.Done()
		}()
	}

	wg.Wait()
}

func main() {
	tmp := []int{9, 3, 2, 5, 8, 4, 1, 6, 7, 0}
	data := tmp[:]

	fmt.Println(data)
	quickSort(data, 0, 9)
	fmt.Println(data)
}
