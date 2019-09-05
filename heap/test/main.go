package test

type Heap struct {
	Size int
	Data [100]int
}

// 插入
func (h *Heap)Push(value int) {
	if h.isFull() {
		return
	}
	h.Data[h.Size] = value
	h.Size++
	h.moveUp()
}

// 取出堆顶元素
func (h *Heap)Pop() int {
	if h.isEmpty() {
		return -999
	}
	min := h.Data[0]
	h.Data[0] = h.Data[h.Size - 1]
	h.Size--
	h.moveDown()
	return min
}

// moveUp
func (h *Heap)moveUp() {
	if h.Size <= 1 {
		return
	}

	curIndex := h.Size - 1
	parentIndex := (curIndex - 1) / 2

	for {
		if h.Data[parentIndex] > h.Data[curIndex] {
			h.Data[parentIndex], h.Data[curIndex] = h.Data[curIndex], h.Data[parentIndex]
		}
		if parentIndex == 0 {
			return
		}
		curIndex = parentIndex
		parentIndex = (curIndex - 1) / 2
	}
}

// moveDown
func (h *Heap)moveDown() {
	if h.Size <= 1 {
		return
	}

	curIndex := 0
	leftChild := curIndex * 2 + 1
	rightChild := curIndex * 2 + 2

	for {
		if h.Data[curIndex] > h.Data[leftChild] && leftChild <= h.Size - 1 {
			h.Data[curIndex], h.Data[leftChild] = h.Data[leftChild], h.Data[curIndex]
			curIndex = leftChild
		} else if h.Data[curIndex] > h.Data[rightChild] && rightChild <= h.Size - 1 {
			h.Data[curIndex], h.Data[rightChild] = h.Data[rightChild], h.Data[curIndex]
			curIndex = rightChild
		} else {
			return
		}
		if curIndex >= h.Size {
			return
		}
		leftChild = curIndex * 2 + 1
		rightChild = curIndex * 2 + 2
	}
}

// 堆是否已满
func (h *Heap) isFull() bool {
	return h.Size == 100
}

// 堆是否为空
func (h *Heap) isEmpty() bool {
	return h.Size == 0
}


