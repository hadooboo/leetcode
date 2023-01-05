import "sort"

func minimumSize(nums []int, maxOperations int) int {
	sort.Ints(nums)
	return sort.Search(10e9, func(i int) bool {
		if i == 0 {
			return false
		}
		c := 0
		for j := len(nums) - 1; j >= 0; j-- {
			if i >= nums[j] {
				break
			}
			c += (nums[j]+i-1)/i - 1
		}
		return c <= maxOperations
	})
}