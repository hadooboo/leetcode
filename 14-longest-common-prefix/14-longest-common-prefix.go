func longestCommonPrefix(strs []string) string {
	m := len(strs[0])
	for _, v := range strs {
		m = min(m, getCommonPrefixLength(strs[0], v))
	}
	return strs[0][:m]
}

func getCommonPrefixLength(s1, s2 string) int {
	for i := 0; ; i++ {
		if i >= len(s1) || i >= len(s2) {
			return i
		}
		if s1[i] != s2[i] {
			return i
		}
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}