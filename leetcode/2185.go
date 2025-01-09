package leetcode

func prefixCount(words []string, pref string) int {
	counter := 0

	for _, word := range words {
		if len(word) < len(pref) {
			continue
		}
		if word[:len(pref)] == pref {
			counter++
		}
	}
	return counter
}
