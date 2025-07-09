package leetcode2

var costMap = map[byte]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

func romanToInt(s string) int {
	acc := 0
	lastGreatest := 0

	for i := len(s) - 1; i >= 0; i-- {
		cost := costMap[s[i]]
		if cost < lastGreatest {
			acc -= cost
		} else {
			lastGreatest = cost
			acc += cost
		}
	}

	return acc
}
