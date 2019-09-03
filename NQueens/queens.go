package NQueens

var Count int = 0

func NQueensDFS(n uint, row uint, col, pie, na int) {
	if row >= n {
		Count++
		return
	}

	// 得倒所有可以放置皇后的位置
	bits := (^(col | pie | na)) & ((1 << n) - 1)

	for bits > 0 {
		// 得倒最后一个空位
		p := bits & (-bits)
		NQueensDFS(n, row + 1, col | p, (pie | p) << 1, (na | p) >> 1)
		bits = bits & (bits - 1)
	}

}
