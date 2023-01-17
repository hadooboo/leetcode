func minFlipsMonoIncr(s string) int {
	cnt, c := make([]int, len(s)+1), 0
	for i, v := range s {
		if v == '0' {
			c++
		}
		cnt[i+1] = c
	}

	dp := make([]int, len(s)+1)
	for i := len(s) - 1; i >= 0; i-- {
		switch s[i] {
		case '0':
			dp[i] = dp[i+1]
		case '1':
			dp[i] = min(dp[i+1]+1, cnt[len(cnt)-1]-cnt[i])
		}
	}

	return dp[0]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}