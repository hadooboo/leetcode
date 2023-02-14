func addBinary(a string, b string) string {
	la, lb := len(a), len(b)
	n := max(la, lb) + 1
	res := make([]byte, 0, n)
	var a1, b1, carry bool
	for i := 0; i < n-1; i++ {
		a1 = i < la && a[la-1-i] == '1'
		b1 = i < lb && b[lb-1-i] == '1'
		switch {
		case a1 && b1 && carry:
			carry = true
			res = append(res, '1')
		case a1 && b1 && !carry,
			a1 && !b1 && carry,
			!a1 && b1 && carry:
			carry = true
			res = append(res, '0')
		case a1 && !b1 && !carry,
			!a1 && b1 && !carry,
			!a1 && !b1 && carry:
			carry = false
			res = append(res, '1')
		case !a1 && !b1 && !carry:
			carry = false
			res = append(res, '0')
		}
	}
	if carry {
		res = append(res, '1')
	}
	return string(reverse(res))
}

func reverse(s []byte) []byte {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
