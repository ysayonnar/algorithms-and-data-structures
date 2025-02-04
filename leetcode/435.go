package leetcode

import (
	"sort"
)

type Intervals [][]int

func (ints Intervals) Len() int {
	return len(ints)
}

func (ints Intervals) Less(i, j int) bool {
	return ints[i][0] < ints[j][0]
}

func (ints Intervals) Swap(i, j int) {
	ints[i], ints[j] = ints[j], ints[i]
}

func eraseOverlapIntervals(intervals [][]int) int {
	sort.Sort(Intervals(intervals))
	ctr := 0
	lastPoint := intervals[0][1]

	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] < lastPoint {
			if intervals[i][1] < lastPoint {
				lastPoint = intervals[i][1]
			}
			ctr++
		} else {
			lastPoint = intervals[i][1]
		}

	}

	return ctr
}
