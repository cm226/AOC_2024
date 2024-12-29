package day19

import (
	util "aox_2024/src/utils"
	"fmt"
	"strings"
)

func parseInput(fileName string) ([]string, [][]string) {

	input := util.FileToMatrix(fileName, ",", util.NoOpConverter)
	towls := input[0]
	target := input[2:]

	return towls, target
}

func possible(target string, towels map[string]bool, cache map[string]int, maxlen int) int {
	if possible, ok := cache[target]; ok {
		return possible
	}

	if len(target) == 0 {
		return 0
	}

	if _, ok := towels[target]; ok {
		cache[target]++
	}

	for i := 1; i < maxlen; i++ {
		if len(target) > i {
			if _, ok := towels[target[0:i]]; ok {
				cache[target] += possible(target[i:], towels, cache, maxlen)
			}
		}
	}

	count := cache[target]

	return count
}

func Day19() {
	towels, targets := parseInput("input.txt")

	maxlen := 0
	towelsMap := map[string]bool{}
	for _, t := range towels {
		towelsMap[t] = true
		if len(t) > maxlen {
			maxlen = len(t)
		}
	}

	validCount := 0
	totalPossible := 0

	cache := map[string]int{}
	for _, target := range targets {
		targetStr := strings.Join(target, "")
		numPerms := possible(targetStr, towelsMap, cache, maxlen+1)

		if numPerms != 0 {
			validCount++
		}
		totalPossible += numPerms
	}

	fmt.Println(validCount)
	fmt.Println(totalPossible)
}
