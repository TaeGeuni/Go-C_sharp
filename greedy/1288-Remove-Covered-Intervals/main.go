package main

import (
	"sort"
)

func removeCoveredIntervals(intervals [][]int) int {
	res := 1

	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] == intervals[j][0] {
			return intervals[i][1] > intervals[j][1]
		}
		return intervals[i][0] < intervals[j][0]
	})
	max := intervals[0][1]

	for i := 1; i < len(intervals); i++ {
		if max < intervals[i][1] {
			max = intervals[i][1]
			res++
		}
	}

	return res
}
