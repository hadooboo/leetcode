import (
	"math"
	"sort"
)

func minimumSize(nums []int, maxOperations int) int {
	sort.Ints(nums)
	max, sum := 0, 0
	for _, v := range nums {
		if max < v {
			max = v
		}
		sum += v
	}
	min := int(math.Ceil(float64(sum) / float64(maxOperations+len(nums))))
	return sort.Search(max-min+1, func(i int) bool {
		i += min
		c := 0
		for j := len(nums) - 1; j >= 0; j-- {
			if i >= nums[j] {
				break
			}
			c += (nums[j]+i-1)/i - 1
		}
		return c <= maxOperations
	}) + min
}