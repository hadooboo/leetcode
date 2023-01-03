var digitToLetterMap = map[byte][]byte{
	'2': {'a', 'b', 'c'},
	'3': {'d', 'e', 'f'},
	'4': {'g', 'h', 'i'},
	'5': {'j', 'k', 'l'},
	'6': {'m', 'n', 'o'},
	'7': {'p', 'q', 'r', 's'},
	'8': {'t', 'u', 'v'},
	'9': {'w', 'x', 'y', 'z'},
}

var res []string

func letterCombinations(digits string) []string {
	res = make([]string, 0)
	doBFS(make([]byte, len(digits)), digits, 0)
	return res
}

func doBFS(arr []byte, digits string, depth int) {
	if depth == len(digits) {
		if len(arr) > 0 {
			res = append(res, string(arr))
		}
		return
	}

	for _, v := range digitToLetterMap[digits[depth]] {
		arr[depth] = v
		doBFS(arr, digits, depth+1)
	}
}