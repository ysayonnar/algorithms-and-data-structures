package leetcode

import (
	"strconv"
)

func isNumber(s byte) bool {
	if s == '1' || s == '2' || s == '3' || s == '4' || s == '5' || s == '6' || s == '7' || s == '8' || s == '9' || s == '0' {
		return true
	}
	return false
}

func decodeString(s string) string {
	result := []byte{}
	leftIdx := 0
	rightIdx := len(s) - 1
	timesBytes := []byte{}

	for i := 0; i < len(s); i++ {
		if s[i] == '[' {
			leftIdx = i
			boundsCounter := 0
			for j := i + 1; j < len(s); j++ {
				if s[j] == '[' {
					boundsCounter++
				} else if s[j] == ']' && boundsCounter == 0 {
					rightIdx = j
					j = len(s)
				} else if s[j] == ']' {
					boundsCounter--
				}
			}
			times, _ := strconv.Atoi(string(timesBytes))
			timesBytes = []byte{}
			temp := []byte(decodeString(s[leftIdx+1 : rightIdx]))
			for i := 0; i < times; i++ {
				result = append(result, temp...)
			}
			i = rightIdx
		} else if isNumber(s[i]) {
			timesBytes = append(timesBytes, s[i])
		} else {
			result = append(result, s[i])
		}
	}

	return string(result)
}
