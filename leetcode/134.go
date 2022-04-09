package leetcode

import "math"

/*
	加油站
*/

func canCompleteCircuit(gas []int, cost []int) int {
	var spare, minIdx int
	minSpare := math.MaxInt64
	for i := 0; i < len(gas); i++ {
		spare += gas[i] - cost[i]
		if spare < minSpare {
			minSpare = spare
			minIdx = i
		}
	}
	if spare < 0 {
		return -1
	}
	return (minIdx + 1) % len(gas)
}
