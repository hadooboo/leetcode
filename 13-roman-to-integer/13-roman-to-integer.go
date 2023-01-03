import (
	"github.com/emirpasic/gods/sets/hashset"
)

var (
	roman = map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	negCaseSet = hashset.New("IV", "IX", "XL", "XC", "CD", "CM")
)

func romanToInt(s string) int {
	res := 0
	for i, v := range []byte(s) {
		if isNeg(s, i) {
			res -= roman[v]
		} else {
			res += roman[v]
		}
	}
	return res
}

func isNeg(s string, i int) bool {
	if i == len(s)-1 {
		return false
	}
	return negCaseSet.Contains(s[i : i+2])
}