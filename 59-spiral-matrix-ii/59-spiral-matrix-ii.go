var (
	dx = []int{0, 1, 0, -1}
	dy = []int{1, 0, -1, 0}
)

func generateMatrix(n int) [][]int {
	res := make([][]int, n)
	for i := range res {
		res[i] = make([]int, n)
	}

	x, y, p := 0, 0, 1
	n -= 1
	for n > 0 {
		for i := 0; i < 4; i++ {
			for j := 0; j < n; j++ {
				res[x][y] = p
				p += 1
				x, y = x+dx[i], y+dy[i]
			}
		}
		n -= 2
		x, y = x+1, y+1
	}
	if n == 0 {
		res[x][y] = p
	}

	return res
}