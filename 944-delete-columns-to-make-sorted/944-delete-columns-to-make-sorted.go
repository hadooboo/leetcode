func minDeletionSize(strs []string) int {
	length := len(strs[0])
	res := 0
	for i := 0; i < length; i++ {
		if !isSorted(strs, i) {
			res++
		}
	}
	return res
}

func isSorted(strs []string, i int) bool {
	for j := 0; j < len(strs)-1; j++ {
		if strs[j][i] > strs[j+1][i] {
			return false
		}
	}
	return true
}