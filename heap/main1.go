package main

import "fmt"

// minHeap
type Heap struct {
	heap []int
}

func (h *Heap) Len() int {
	return len(h.heap)
}

func (h *Heap) Less(i, j int) bool {
	return h.heap[i] < h.heap[j]
}

func (h *Heap) Swap(i, j int) {
	h.heap[i], h.heap[j] = h.heap[j], h.heap[i]
}

func (h *Heap) Push(x interface{}) {
	h.heap = append(h.heap, x.(int))
}

func (h *Heap) Pop() interface{} {
	v := h.heap[0]
	h.heap = h.heap[1:]
	return v
}

func (h *Heap) List()  {
	fmt.Println(h.heap)
}




