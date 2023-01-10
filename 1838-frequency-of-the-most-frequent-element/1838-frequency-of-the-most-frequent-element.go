import "sort"

func maxFrequency(nums []int, k int) int {
	n := len(nums)

	sort.Ints(nums)

	x, c, res := 0, 0, 1
	for i := 1; i < n; i++ {
		c += (nums[i] - nums[i-1]) * (i - x)
		for c > k && x < i {
			c -= nums[i] - nums[x]
			x++
		}
		max := i - x + 1
		if res < max {
			res = max
		}
	}

	return res
}