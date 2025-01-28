package leetcode

import "sort"

func successfulPairs(spells []int, potions []int, success int64) []int {
	sort.Ints(potions) // сортируем potions для бинарного поиска
	result := make([]int, len(spells))

	for i, spell := range spells {
		idx := len(potions)
		l, r := 0, len(potions)-1
		for l <= r {
			mid := l + (r-l)/2
			if potions[mid]*spell >= int(success) {
				idx = mid
				r = mid - 1
			} else {
				l = mid + 1
			}
		}

		result[i] = len(potions) - idx
	}
	return result
}
