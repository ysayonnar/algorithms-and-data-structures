package leetcode

import "sort"

type Points [][]int

func (p Points) Len() int {
	return len(p)
}

func (p Points) Less(i, j int) bool {
	return p[i][1] < p[j][1]
}

func (p Points) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func findMinArrowShots(points [][]int) int {
	sort.Sort(Points(points))

	ctr := 1
	lastPoint := points[0][1]

	for i := 1; i < len(points); i++ {
		if points[i][0] > lastPoint {
			ctr++
			lastPoint = points[i][1]
		}
	}

	return ctr
}
