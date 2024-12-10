package day10

import (
	util "aox_2024/src/utils"
	"fmt"
	"strconv"
)

func recurseAll(start util.Point, topmap [][]int, peaks *[][]int, part2 bool) {
	current := topmap[start.Y][start.X]

	if current == 9 {
		if part2 {
			(*peaks)[start.Y][start.X] += 1
		} else {

			(*peaks)[start.Y][start.X] = 1
		}
		return
	}

	left := util.Point{X: start.X - 1, Y: start.Y}
	right := util.Point{X: start.X + 1, Y: start.Y}
	up := util.Point{X: start.X, Y: start.Y - 1}
	down := util.Point{X: start.X, Y: start.Y + 1}

	corner := util.Point{Y: len(topmap), X: len(topmap[0])}
	if left.Inside(corner) {
		lValue := util.IndexPoint(&topmap, left)
		if lValue == current+1 {
			recurseAll(left, topmap, peaks, part2)
		}
	}
	if right.Inside(corner) {
		rValue := util.IndexPoint(&topmap, right)
		if rValue == current+1 {
			recurseAll(right, topmap, peaks, part2)
		}
	}
	if up.Inside(corner) {
		uValue := util.IndexPoint(&topmap, up)
		if uValue == current+1 {
			recurseAll(up, topmap, peaks, part2)
		}
	}
	if down.Inside(corner) {
		dValue := util.IndexPoint(&topmap, down)
		if dValue == current+1 {
			recurseAll(down, topmap, peaks, part2)
		}
	}
}

func part1And2(topomap [][]int) {

	totalScores := 0
	for y, line := range topomap {
		for x, height := range line {
			if height == 0 {
				peaks := make([][]int, len(topomap))
				for i, _ := range peaks {
					peaks[i] = make([]int, len(topomap[i]))
				}
				recurseAll(util.Point{X: x, Y: y}, topomap, &peaks, true)

				for _, peakLine := range peaks {
					for _, peak := range peakLine {
						totalScores += peak
					}
				}
			}
		}
	}
	fmt.Println(totalScores)

}

func Day10() {

	topmap := util.FileToMatrix("input.txt", "", func(s string) int {
		i, e := strconv.Atoi(s)
		util.PanicIfError(e)
		return i
	})
	part1And2(topmap)

}
