package day13

import (
	util "aox_2024/src/utils"
	"bufio"
	"os"
	"strconv"
	"strings"
)

type block struct {
	A     util.Point
	B     util.Point
	point util.Point
}

const MaxUint = ^uint(0)
const MinUint = 0
const MaxInt = int(MaxUint >> 1)

func parseBlock(lines []string) block {
	if len(lines) != 3 {
		panic("failed to parse")
	}

	parseLine := func(line string) util.Point {
		movesS := strings.Split(line, ":")[1]
		xys := strings.Split(movesS, ",")
		x, e := strconv.Atoi(strings.TrimSpace(xys[0])[2:])
		util.PanicIfError(e)
		y, e := strconv.Atoi(strings.TrimSpace(xys[1])[2:])
		util.PanicIfError(e)

		return util.Point{X: x, Y: y}
	}

	return block{
		A:     parseLine(lines[0]),
		B:     parseLine(lines[1]),
		point: parseLine(lines[2]),
	}
}

func parseInput(filename string) []block {

	inputfile, error := os.Open(filename)
	util.PanicIfError(error)

	scanner := bufio.NewScanner(inputfile)

	blocks := []block{}
	blockLines := []string{}
	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			block := parseBlock(blockLines)
			blocks = append(blocks, block)
			blockLines = make([]string, 0)
		} else {
			blockLines = append(blockLines, line)
		}
	}
	block := parseBlock(blockLines)
	blocks = append(blocks, block)

	return blocks
}

func findFactors(block block) int {

	position := util.Point{}
	mincost := MaxInt
	//	a := 0
	//	for (a*block.A.X < block.point.X) && (a*block.A.Y < block.point.Y) {
	for a := range 102 {
		if a == 101 {
			if mincost == MaxInt {
				mincost = 0
			}
			continue
		}
		//b := 0
		//for (b*block.B.X < block.point.X) && (b*block.B.Y < block.point.Y) {
		for b := range 101 {

			position.X = a * block.A.X
			position.X += b * block.B.X

			position.Y = a * block.A.Y
			position.Y += b * block.B.Y

			cost := (a * 3) + b
			if position == block.point && cost < mincost {
				mincost = cost
			}
		}
	}
	return mincost
}

func Day13() {
	blocks := parseInput("input.txt")

	mincost := 0

	for _, block := range blocks {
		mincost += findFactors(block)
	}

	println(mincost)

}
