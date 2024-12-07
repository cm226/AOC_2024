package day2

import (
	util "aox_2024/src/utils"
	"fmt"
	"strconv"
)

func isSafe(report []int) bool {
	direction := true // true == ascending
	for i := range len(report) - 1 {
		cur := report[i]
		next := report[i+1]

		if i == 0 {
			direction = cur < next
		}

		directionCheck := (cur < next) == direction
		absDif := util.Abs(cur, next)

		if !(directionCheck && absDif >= 1 && absDif <= 3) {
			return false
		}
	}
	return true
}

func part1(reports [][]int) {
	safeCount := 0
	for _, report := range reports {
		if isSafe(report) {
			safeCount += 1
		}
	}
	fmt.Println(safeCount)

}

func part2(reports [][]int) {
	safeCount := 0
	for _, report := range reports {
		safe := isSafe(report)

		if !safe {
			for i := range len(report) {
				newReport := make([]int, len(report))
				copy(newReport, report[:i])
				copy(newReport[i:], report[i+1:])
				newReport = newReport[:len(newReport)-1]
				if isSafe(newReport) {
					safeCount += 1
					break
				}
			}
		} else {
			safeCount += 1
		}
	}
	fmt.Println(safeCount)

}
func Day2() {

	reports := util.FileToMatrix("day2.txt", " ",
		func(s string) int {
			i, e := strconv.Atoi(s)
			util.PanicIfError(e)
			return i
		})

	part2(reports)

}
