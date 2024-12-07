package main

import (
	util "aox_2024/src/utils"
	"fmt"
	"slices"
	"strings"
)

type visited struct {
	x      int
	y      int
	orient string
}

func getDirection(guard string) (int, int) {
	switch guard {
	case ">":
		return 1, 0
	case "<":
		return -1, 0
	case "^":
		return 0, -1
	case "V":
		return 0, 1
	}
	panic("invalid guard")
}

func getGuardSymbol(from util.Point, to util.Point) string {

	guard := ">"
	if from.X > to.X {
		guard = "<"
	} else if from.Y > to.Y {
		guard = "^"
	} else if from.Y < to.Y {
		guard = "V"
	}
	return guard
}

func rotSymbol(cur string) string {
	switch cur {
	case ">":
		return "V"
	case "V":
		return "<"
	case "<":
		return "^"
	case "^":
		return ">"
	}
	panic("unknown guard")
}

func moveGuard(board *[][]string, from util.Point, to util.Point) {

	guard := getGuardSymbol(from, to)
	(*board)[from.Y][from.X] = "X"
	(*board)[to.Y][to.X] = guard
}

func findGuard(board *[][]string) (int, int) {

	row := -1
	column := -1
	for i, line := range *board {
		for j, c := range line {
			if strings.Contains("V<>^", c) {
				row = i
				column = j
				break
			}
		}
		if row != -1 && column != -1 {
			break
		}
	}
	return row, column
}

func simStep(board *[][]string, gaurdY int, gaurdX int) visited {
	dirx, diry := getDirection((*board)[gaurdY][gaurdX])

	nextGuardPosRow := gaurdY + diry
	nextGuardPosCol := gaurdX + dirx

	if nextGuardPosCol < 0 || nextGuardPosRow < 0 || nextGuardPosRow >= len(*board) || nextGuardPosCol >= len((*board)[nextGuardPosRow]) {
		return visited{
			x: nextGuardPosCol,
			y: nextGuardPosRow,
		}
	}

	guard := (*board)[gaurdY][gaurdX]
	for !((*board)[nextGuardPosRow][nextGuardPosCol] == "." || (*board)[nextGuardPosRow][nextGuardPosCol] == "X") {
		guard = rotSymbol(guard)
		dirx, diry = getDirection(guard)
		nextGuardPosRow = gaurdY + diry
		nextGuardPosCol = gaurdX + dirx
	}

	moveGuard(board, util.Point{
		X: gaurdX,
		Y: gaurdY,
	},
		util.Point{
			X: nextGuardPosCol,
			Y: nextGuardPosRow,
		})

	return visited{
		x:      nextGuardPosCol,
		y:      nextGuardPosRow,
		orient: (*board)[nextGuardPosRow][nextGuardPosCol],
	}
}

func hasLooped(visited []visited) bool {

	if len(visited) < 2 {
		return false
	}
	last := visited[len(visited)-1]

	for i := range len(visited) - 1 {
		if last.x == visited[i].x && last.y == visited[i].y && last.orient == visited[i].orient {
			return true
		}
	}
	return false
}

func day6Part1(board [][]string) []visited {

	// find guard
	gaurdY, gaurdX := findGuard(&board)

	if gaurdY == -1 || gaurdX == -1 {
		panic("failed to find guard")
	}

	vpos := []visited{}
	for !hasLooped(vpos) {
		nextVisit := simStep(&board, gaurdY, gaurdX)
		if nextVisit.y >= len(board) || nextVisit.x >= len(board[0]) {
			break
		}
		gaurdX = nextVisit.x
		gaurdY = nextVisit.y
		vpos = append(vpos, nextVisit)
	}

	util.PrintMatrix(board)

	visited := 0
	for _, row := range board {
		for _, c := range row {
			if c == "X" {
				visited += 1
			}
		}
	}
	fmt.Println(visited + 1)
	return vpos
}

func getObsicleLocations(v []visited) []util.Point {
	locaitions := []util.Point{}

	compare := func(p1 util.Point) func(util.Point) bool {
		return func(p2 util.Point) bool {
			return p1.X == p2.X && p1.Y == p2.Y
		}
	}
	for _, pos := range v {
		pt := util.Point{
			X: pos.x,
			Y: pos.y,
		}
		if !slices.ContainsFunc(locaitions, compare(pt)) {
			locaitions = append(locaitions, pt)
		}
	}
	return locaitions
}

func day6Part2(origboard [][]string, locations []util.Point) int {

	loopCount := 0
	board := make([][]string, len(origboard))

	for i := range board {
		board[i] = make([]string, len(origboard[i]))
		copy(board[i], origboard[i])
	}
	gaurdOrigRow, gaurdOrigCol := findGuard(&origboard)

	for _, loc := range locations {
		i := loc.Y
		j := loc.X

		if i < 0 || j < 0 || i > len(board)-1 || j > len(board[i])-1 {
			continue
		}

		gposRow := gaurdOrigRow
		gposCol := gaurdOrigCol
		if gposRow == i && gposCol == j {
			continue
		}

		if board[i][j] == "#" {
			continue
		}
		// insert obstical
		board[i][j] = "#"

		vpos := []visited{}
		for !hasLooped(vpos) {
			nextVisit := simStep(&board, gposRow, gposCol)
			gposRow = nextVisit.y
			gposCol = nextVisit.x
			if nextVisit.y >= len(board) || nextVisit.x >= len(board[0]) || nextVisit.x < 0 || nextVisit.y < 0 {
				break
			}
			vpos = append(vpos, nextVisit)
		}
		if hasLooped(vpos) {
			loopCount += 1
		}

		for i := range board {
			copy(board[i], origboard[i])
		}
	}
	return loopCount
}

func day6() {
	origboard := util.FileToMatrix("input.txt", "", util.NoOpConverter)
	board := make([][]string, len(origboard))
	for i := range board {
		board[i] = make([]string, len(origboard[i]))
		copy(board[i], origboard[i])
	}
	visited := day6Part1(board)
	locations := getObsicleLocations(visited)

	ch := make(chan int)
	splits := 23
	chunks := len(locations) / splits

	for i := range splits + 1 {
		go func() {
			fmt.Println(i*chunks, (i+1)*chunks)
			ch <- day6Part2(origboard, locations[i*chunks:(i+1)*chunks])
		}()
	}

	chunkResults := []int{}
	for range splits + 1 {
		x := <-ch // Receiving the value from the channel
		chunkResults = append(chunkResults, x)
	}

	total := 0
	for _, c := range chunkResults {
		total += c
	}
	fmt.Println(total)

}
