import (
	"math"
)

func minimumSize(nums []int, maxOperations int) int {
	max, sum := 0, 0
	for _, v := range nums {
		if max < v {
			max = v
		}
		sum += v
	}
	min := int(math.Ceil(float64(sum) / float64(maxOperations+len(nums))))
	return doBinarySearch(max-min+1, func(i int) bool {
		i += min
		c := 0
		for _, v := range nums {
			c += (v+i-1)/i - 1
		}
		return c <= maxOperations
	}) + min
}

func doBinarySearch(n int, f func(int) bool) int {
	i, j := 0, n
	for i < j {
		h := int(uint(i+j) >> 1)
		if !f(h) {
			i = h + 1
		} else {
			j = h
		}
	}
	return i
}