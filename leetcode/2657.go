package leetcode

func findThePrefixCommonArray(a []int, b []int) []int {
	set := make(map[int]int) // map[value]amount
	result := make([]int, len(a))
	ctr := 0

	for i := 0; i < len(a); i++ {
		set[a[i]]++
		set[b[i]]++
		if set[a[i]] > 1 {
			set[a[i]] = 0
			ctr++
		}
		if set[b[i]] > 1 {
			set[b[i]] = 0
			ctr++
		}
		result[i] = ctr
	}
	return result
}

func FindThePrefixCommonArray(a, b []int) []int {
	return findThePrefixCommonArray(a, b)
}
