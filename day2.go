package main

import (
	"fmt"
	"strconv"
)

func isSafe(report []string) bool {
	direction := true // true == ascending
	for i := range len(report) - 1 {
		cur, e := strconv.Atoi(report[i])
		next, e2 := strconv.Atoi(report[i+1])

		panicIfError(e)
		panicIfError(e2)

		if i == 0 {
			direction = cur < next
		}

		directionCheck := (cur < next) == direction
		absDif := abs(cur, next)

		if !(directionCheck && absDif >= 1 && absDif <= 3) {
			return false
		}
	}
	return true
}

func part1(reports [][]string) {
	safeCount := 0
	for _, report := range reports {
		if isSafe(report) {
			safeCount += 1
		}
	}
	fmt.Println(safeCount)

}

func part2(reports [][]string) {
	safeCount := 0
	for _, report := range reports {
		safe := isSafe(report)

		if !safe {
			for i := range len(report) {
				newReport := make([]string, len(report))
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
func day2() {

	reports := fileToMatrix("day2.txt")
	part2(reports)

}
