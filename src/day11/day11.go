package day11

import (
	util "aox_2024/src/utils"
	"container/list"
	"fmt"
	"strconv"
)

func runStone(stone int) []int {

	valStr := strconv.Itoa(stone)
	if stone == 0 {
		return []int{1}
	} else if len(valStr)%2 == 0 {
		stone1Str := valStr[:len(valStr)/2]
		stone2Str := valStr[len(valStr)/2:]

		stone1, e := strconv.Atoi(stone1Str)
		util.PanicIfError(e)
		stone2, e := strconv.Atoi(stone2Str)
		util.PanicIfError(e)

		return []int{stone1, stone2}

	} else {
		return []int{stone * 2024}
	}

}

func runStep(stones *list.List) {
	front := stones.Front()

	for front != nil {
		next := front.Next()
		blinked := runStone(front.Value.(int))
		if len(blinked) == 2 {
			stones.InsertAfter(blinked[1], front)
		}
		front.Value = blinked[0]
		front = next
	}
}

func part1(allStones []int) {

	stoneList := list.New()
	for _, stone := range allStones {
		stoneList.PushBack(stone)
	}

	for range 25 {
		runStep(stoneList)
	}

	fmt.Println(stoneList.Len())
}

func recurseAll(stone int, currentDepth int, targetDepth int, cache *map[int]map[int]int) int {

	numStones := 0
	depth, seenStone := (*cache)[stone]
	if seenStone {
		count, seenDepth := depth[targetDepth-currentDepth]
		if seenDepth {
			return count
		}
	} else {
		(*cache)[stone] = map[int]int{}
	}

	expanded := runStone(stone)

	if currentDepth+1 != targetDepth {
		numStones += recurseAll(expanded[0], currentDepth+1, targetDepth, cache)
		if len(expanded) > 1 {
			numStones += recurseAll(expanded[1], currentDepth+1, targetDepth, cache)
		}
	} else {
		return len(expanded)
	}

	(*cache)[stone][targetDepth-currentDepth] = numStones

	return numStones
}

func part2(allStones []int) {
	cache := map[int]map[int]int{} // stone -> depth -> numstones
	numIter := 75
	allStoneCount := 0

	for _, stone := range allStones {
		recurseAll(stone, 0, numIter, &cache)
		allStoneCount += cache[stone][numIter]
	}
	fmt.Println(allStoneCount)

}

func Day11() {
	allStones := util.FileToSlice("input.txt", " ", func(s string) int {
		i, e := strconv.Atoi(s)
		util.PanicIfError(e)
		return i
	})

	part1(allStones)
	part2(allStones)

}
