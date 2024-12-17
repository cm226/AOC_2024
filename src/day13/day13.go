package day13

import (
	util "aox_2024/src/utils"
	"bufio"
	"fmt"
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

	b := block{
		A:     parseLine(lines[0]),
		B:     parseLine(lines[1]),
		point: parseLine(lines[2]),
	}

	b.point.X += 10000000000000
	b.point.Y += 10000000000000
	return b
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

func solve(b block) (int, int) {
	// solve with liner algebra

	a1 := (b.point.Y * b.A.X) - (b.A.Y * b.point.X)
	a2 := -(b.A.Y * b.B.X) + (b.B.Y * b.A.X)

	if a1%a2 != 0 {
		// no solution
		return 0, 0
	}

	bCount := a1 / a2

	aCount := (b.point.Y - (b.B.Y * bCount))
	if aCount%b.A.Y != 0 {
		// no solution
		return 0, 0
	}
	aCount = aCount / b.A.Y

	return aCount, bCount

}

func Day13() {
	blocks := parseInput("input.txt")

	mincost := 0

	for _, block := range blocks {
		aCount, bCount := solve(block)
		mincost += aCount*3 + bCount
		fmt.Println(aCount, bCount)
	}

	println(mincost)

}
