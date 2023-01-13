func maximumGain(s string, x int, y int) int {
	t1, t2 := "ab", "ba"
	if x < y {
		t1, t2 = t2, t1
		x, y = y, x
	}
	res := 0
	stack, p := make([]byte, len(s)), -1
	for _, v := range []byte(s) {
		if p >= 0 && stack[p] == t1[0] && v == t1[1] {
			p--
			res += x
		} else {
			p++
			stack[p] = v
		}
	}
	s, p = string(stack[:p+1]), -1
	for _, v := range []byte(s) {
		if p >= 0 && stack[p] == t2[0] && v == t2[1] {
			p--
			res += y
		} else {
			p++
			stack[p] = v
		}
	}

	return res
}