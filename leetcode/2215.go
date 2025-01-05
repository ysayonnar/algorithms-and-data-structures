package leetcode

func difference(map1, map2 map[int]bool) []int {
	diff := []int{}
	for value, _ := range map1 {
		if _, ok := map2[value]; !ok {
			diff = append(diff, value)
		}
	}
	return diff
}

func findDifference(nums1 []int, nums2 []int) [][]int {
	map1 := make(map[int]bool)
	map2 := make(map[int]bool)

	for _, num := range nums1 {
		map1[num] = true
	}

	for _, num := range nums2 {
		map2[num] = true
	}

	return [][]int{difference(map1, map2), difference(map2, map1)}
}
