func lengthOfLIS(nums []int) int {
	res := make([]int, 0)

	for _, v := range nums {
		idx := sort.Search(len(res), func(i int) bool {
			return res[i] >= v
		})
		if idx == len(res) {
			res = append(res, v)
		} else {
			res[idx] = v
		}
	}

	return len(res)
}