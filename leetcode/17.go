package leetcode

func letterCombinations(digits string) []string {
	result := []string{}
	if len(digits) == 0 {
		return result
	}

	symbols := map[byte][]byte{
		'2': []byte{'a', 'b', 'c'},
		'3': []byte{'d', 'e', 'f'},
		'4': []byte{'g', 'h', 'i'},
		'5': []byte{'j', 'k', 'l'},
		'6': []byte{'m', 'n', 'o'},
		'7': []byte{'p', 'q', 'r', 's'},
		'8': []byte{'t', 'u', 'v'},
		'9': []byte{'w', 'x', 'y', 'z'},
	}

	var backtrack func(index int, combintaion string)
	backtrack = func(index int, combination string) {
		if index == len(digits) {
			result = append(result, combination)
			return
		}

		letters := symbols[digits[index]]
		for i := 0; i < len(letters); i++ {
			backtrack(index+1, combination+string(letters[i]))
		}
	}

	backtrack(0, "")
	return result
}
