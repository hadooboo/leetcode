func detectCapitalUse(word string) bool {
	head, tail := word[0], []byte(word[1:])
	if isUpper(head) {
		return isAllLower(tail) || isAllUpper(tail)
	} else {
		return isAllLower(tail)
	}
}

func isAllUpper(word []byte) bool {
	return all(word, isUpper)
}

func isAllLower(word []byte) bool {
	return all(word, isLower)
}

func all(word []byte, f func(c byte) bool) bool {
	for _, c := range []byte(word) {
		if !f(c) {
			return false
		}
	}
	return true
}

func isUpper(c byte) bool {
	return c >= 'A' && c <= 'Z'
}

func isLower(c byte) bool {
	return c >= 'a' && c <= 'z'
}