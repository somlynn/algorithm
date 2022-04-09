package leetcode

import (
	"testing"
)

func TestCanCompleteCircuit(t *testing.T) {
	gas := []int{1, 2, 3, 4, 5}
	cost := []int{3, 4, 5, 1, 2}
	canCompleteCircuit(gas, cost)
}
