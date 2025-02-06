package leetcode

func isPalindrome123(x int) bool {
	//НОРМАЛЬНОЕ РЕШЕНИЕ
	if x < 0 {
		return false
	} else if x < 10 {
		return true
	}
	digits := []int{}
	for x >= 10 {
		remainder := x % 10
		digits = append(digits, remainder)
		x /= 10
	}
	digits = append(digits, x)

	i, j := 0, len(digits)-1
	for i <= j {
		if digits[i] != digits[j] {
			return false
		}
		i++
		j--
	}
	return true

	//ПИДОРСКОЕ РЕШЕНИЕ!!!!
	// str := strconv.Itoa(x)
	// i, j := 0, len(str)-1
	// for i <= j {
	// 	if str[i] != str[j] {
	// 		return false
	// 	}
	// 	i++
	// 	j--
	// }
	// return true
}
