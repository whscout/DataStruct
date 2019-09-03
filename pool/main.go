package main

import (
	"fmt"
	"sync"
	"time"
)

func Gopool() {
	start := time.Now()
	wg := new(sync.WaitGroup)
	data := make(chan int, 100)

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			for _ = range data {
				//fmt.Println("goroutine:", n, i, v)
			}
		}(i)
	}

	for i := 0; i < 200; i++ {
		data <- i
	}
	close(data)
	wg.Wait()
	end := time.Now()
	fmt.Println(end.Sub(start))
}

func main() {
	Gopool()
}