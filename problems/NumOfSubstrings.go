package problems

import (
	"strings"
)

var MOD float64 = 1e9 + 7
var DICT map[string]int = map[string]int{}

func NumSub(s string) int {
	arr := strings.Split(s, "0")
	output := 0
	for _, subgroups := range arr {
		output += sumTotals(subgroups)
	}
	return output % int(MOD)
}

func sumTotals(s string) int {
	if DICT[s] != 0 {
		return DICT[s]
	}

	n := strings.Count(s, "1")
	sum := 0
	for i := 1; i <= n; i++ {
		sum += i
	}

	return sum
}
