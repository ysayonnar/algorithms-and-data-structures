package leetcode

func closeStrings(word1 string, word2 string) bool {
	symbols1 := make(map[byte]int)
	symbols2 := make(map[byte]int)
	occurrences1 := make(map[int]int)
	occurrences2 := make(map[int]int)

	for i := 0; i < len(word1); i++ {
		symbols1[word1[i]]++
	}

	for i := 0; i < len(word2); i++ {
		symbols2[word2[i]]++
	}

	if len(symbols1) != len(symbols2) {
		return false
	}

	for k1, v1 := range symbols1 {
		if _, ok := symbols2[k1]; !ok {
			return false
		}
		occurrences1[v1]++
	}

	for k2, v2 := range symbols2 {
		if _, ok := symbols1[k2]; !ok {
			return false
		}
		occurrences2[v2]++
	}

	if len(occurrences1) != len(occurrences2) {
		return false
	}

	for k := range occurrences1 {
		if _, ok := occurrences2[k]; !ok {
			return false
		}

		if occurrences1[k] != occurrences2[k] {
			return false
		}
	}

	return true
}
