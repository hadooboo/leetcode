package main

import "fmt"

func ambiguousCoordinates(s string) []string {
	res := make([]string, 0)

	str := s[1 : len(s)-1]
	for i := 1; i < len(str); i++ {
		s1, s2 := str[:i], str[i:]
		for _, v := range findPossibleCoordinates(s1) {
			for _, w := range findPossibleCoordinates(s2) {
				res = append(res, fmt.Sprintf("(%v, %v)", v, w))
			}
		}
	}

	return res
}

func findPossibleCoordinates(s string) []string {
	res := make([]string, 0)

	for i := 1; i <= len(s); i++ {
		if i >= 2 && s[0] == '0' {
			continue
		}
		if i < len(s) && s[len(s)-1] == '0' {
			continue
		}
		v := s
		if i < len(s) {
			v = fmt.Sprintf("%v.%v", s[:i], s[i:])
		}
		res = append(res, v)
	}

	return res
}
