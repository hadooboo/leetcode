func jump(nums []int) int {
	dp := make([]int, len(nums))
	for i := range dp {
		dp[i] = math.MaxInt
	}
    dp[0] = 0

	for i, v := range nums {
		for j := 1; j <= v && i+j < len(nums); j++ {
			dp[i+j] = min(dp[i]+1, dp[i+j])
		}
	}

	return dp[len(nums)-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}