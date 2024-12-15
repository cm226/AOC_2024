package day14

import (
	util "aox_2024/src/utils"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

type robotData struct {
	pos util.Point
	vel util.Point
}

func parseInput(input string) []robotData {
	posVel := util.FileToMatrix(input, " ", util.NoOpConverter)

	robots := []robotData{}

	for _, line := range posVel {
		if len(line) != 2 {
			panic("Failed to parse ")
		}

		positionS := line[0]
		velS := line[1]

		positionS = positionS[2:]
		velS = velS[2:]

		posss := strings.Split(positionS, ",")
		velss := strings.Split(velS, ",")

		p1, e := strconv.Atoi(posss[0])
		util.PanicIfError(e)
		p2, e := strconv.Atoi(posss[1])
		util.PanicIfError(e)

		v1, e := strconv.Atoi(velss[0])
		v2, e := strconv.Atoi(velss[1])

		data := robotData{
			pos: util.Point{
				X: p1,
				Y: p2,
			},
			vel: util.Point{
				X: v1,
				Y: v2,
			},
		}

		robots = append(robots, data)
	}
	return robots

}

func stepRobot(robot *robotData, spacex int, spacey int) {

	newPos := (*robot).pos.Add((*robot).vel)

	if (*robot).vel.X > 0 {
		newPos.X = newPos.X % spacex
	}

	if (*robot).vel.Y > 0 {
		newPos.Y = newPos.Y % spacey
	}

	if (*robot).vel.X < 0 {
		if newPos.X < 0 {
			newPos.X = spacex + newPos.X
		}
	}

	if (*robot).vel.Y < 0 {
		if newPos.Y < 0 {
			newPos.Y = spacey + newPos.Y
		}
	}

	(*robot).pos = newPos
}

func quadrentCount(robots []robotData, tl util.Point, br util.Point) int {

	count := 0
	for _, robot := range robots {
		if robot.pos.X <= br.X && robot.pos.X >= tl.X &&
			robot.pos.Y <= br.Y && robot.pos.Y >= tl.Y {
			count++
		}
	}
	return count

}

func calcQuadrents(robots []robotData, boardX int, boardY int) []int {

	q1Count := quadrentCount(robots,
		util.Point{
			X: 0,
			Y: 0,
		},
		util.Point{
			X: (boardX / 2) - 1,
			Y: (boardY / 2) - 1,
		})

	q2Count := quadrentCount(robots,
		util.Point{
			X: (boardX / 2) + 1,
			Y: 0,
		},
		util.Point{
			X: boardX - 1,
			Y: (boardY / 2) - 1,
		})

	q3Count := quadrentCount(robots,
		util.Point{
			X: 0,
			Y: (boardY / 2) + 1,
		},
		util.Point{
			X: (boardX / 2) - 1,
			Y: boardY - 1,
		})
	q4Count := quadrentCount(robots,
		util.Point{
			X: (boardX / 2) + 1,
			Y: (boardY / 2) + 1,
		},
		util.Point{
			X: boardX - 1,
			Y: boardY - 1,
		})

	return []int{q1Count, q2Count, q3Count, q4Count}
}

func part1(robots []robotData) {

	boardX := 101
	boardY := 103
	seconds := 100

	for range seconds {
		fmt.Println("")
		// print board
		board := make([][]string, boardY)
		for i, c := range board {
			c = make([]string, boardX)

			for j := range c {
				c[j] = "."
				for _, robot := range robots {
					if robot.pos.X == j && robot.pos.Y == i {
						c[j] = "#"
					}
				}
			}
			//fmt.Println(c)
		}

		for i := range robots {
			stepRobot(&robots[i], boardX, boardY)
		}

	}

	quadrents := calcQuadrents(robots, boardX, boardY)
	fmt.Println(quadrents[0] * quadrents[1] * quadrents[2] * quadrents[3])
}

func seeTree(robots []robotData) bool {
	lines := [][]util.Point{}

	slices.SortFunc(robots, func(a, b robotData) int {
		return a.pos.X - b.pos.X
	})

	for robot := range robots {
		for linei := range lines {
			if lines[linei][len(lines[linei])-1].Y == robots[robot].pos.Y && lines[linei][len(lines[linei])-1].X == robots[robot].pos.X-1 {
				lines[linei] = append(lines[linei], robots[robot].pos)
				if len(lines[linei]) >= 16 {
					return true
				}
				break
			}
		}
		lines = append(lines, []util.Point{robots[robot].pos})
	}
	return false
}

func part2(robots []robotData) {

	boardX := 101
	boardY := 103
	seconds := 10000

	for second := range seconds {
		fmt.Println("")
		for i := range robots {
			stepRobot(&robots[i], boardX, boardY)
		}
		if seeTree(robots) {
			fmt.Println(second + 1)

			// print board
			board := make([][]string, boardY)
			for i, c := range board {
				c = make([]string, boardX)

				for j := range c {
					c[j] = "."
					for _, robot := range robots {
						if robot.pos.X == j && robot.pos.Y == i {
							c[j] = "#"
						}
					}
				}
				fmt.Println(c)
			}

			break
		}

	}
}

func Day14() {
	robots := parseInput("input.txt")
	//part1(robots)
	part2(robots)
}
