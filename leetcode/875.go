package leetcode

func minEatingSpeed(piles []int, h int) int {
	minimal := 1
	maximum := piles[0]

	for _, pile := range piles {
		if pile > maximum {
			maximum = pile
		}
	}

	k := 0
	for minimal <= maximum {
		mid := minimal + (maximum-minimal)/2
		currentHours := 0

		for _, pile := range piles {
			hour := pile / mid
			if hour*mid < pile {
				hour++
			}
			currentHours += hour
		}

		if currentHours > h {
			minimal = mid + 1
		} else {
			k = mid
			maximum = mid - 1
		}
	}
	return k
}
