import "sort"

func minimumSize(nums []int, maxOperations int) int {
	return sort.Search(10e9, func(i int) bool {
		if i == 0 {
			return false
		}
		c := 0
		for _, v := range nums {
			c += (v+i-1)/i - 1
		}
		return c <= maxOperations
	})
}