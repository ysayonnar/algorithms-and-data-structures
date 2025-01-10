package leetcode

func countChars(word string) map[byte]int {
	count := make(map[byte]int)
	for i := 0; i < len(word); i++ {
		count[word[i]]++
	}
	return count
}

func wordSubsets(words1 []string, words2 []string) []string {
	words2Count := make(map[byte]int)
	for _, word := range words2 {
		wordCount := countChars(word)
		for k, v := range wordCount {
			if v > words2Count[k] {
				words2Count[k] = v
			}
		}
	}

	result := []string{}

	for _, word := range words1 {
		wordCount := countChars(word)
		isUniversal := true

		for k, v := range words2Count {
			if wordCount[k] < v {
				isUniversal = false
				break
			}
		}

		if isUniversal {
			result = append(result, word)
		}
	}

	return result
}
