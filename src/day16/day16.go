package day16

import (
	util "aox_2024/src/utils"
	"fmt"
	"slices"
)

type node struct {
	pos       util.Point
	direction int
	score     int
}

func listContains(list []node, n node) bool {

	sFn := func(snode node) bool {
		if snode.pos == n.pos {
			return true
		}
		return false
	}
	if slices.ContainsFunc(list, sFn) {
		return true
	}
	return false
}

func listIndex(list []node, n node) int {

	iFn := func(sNode node) bool {
		if sNode.pos == n.pos {
			return true
		}
		return false
	}
	return slices.IndexFunc(list, iFn)
}

func aStar(track [][]string, start util.Point, end util.Point) int {

	open := []node{
		node{
			pos:       start,
			direction: 1,
			score:     0,
		},
	}
	closed := []node{}

	for len(open) != 0 {
		lowest := open[0].score
		currentNode := open[0]
		openi := 0
		for i, n := range open {
			if n.score < lowest {
				currentNode = n
				openi = i
			}
		}

		open = slices.Delete(open, openi, openi+1)
		closed = append(closed, currentNode)

		if currentNode.pos == end {
			return currentNode.score
		}

		// adjacents
		nilNode := node{
			score: -1,
		}
		calcScore := func(n1 node, n2 node) int {
			turns := util.Abs(n1.direction, n2.direction)
			if turns == 3 {
				turns = 1
			}
			return (turns * 1000) + 1 + currentNode.score
		}
		up := node{
			pos:       currentNode.pos.Add(util.PUp()),
			direction: 0,
		}
		down := node{
			pos:       currentNode.pos.Add(util.PDown()),
			direction: 2,
		}
		right := node{
			pos:       currentNode.pos.Add(util.PRight()),
			direction: 1,
		}
		left := node{
			pos:       currentNode.pos.Add(util.PLeft()),
			direction: 3,
		}
		up.score = calcScore(up, currentNode)
		down.score = calcScore(down, currentNode)
		left.score = calcScore(left, currentNode)
		right.score = calcScore(right, currentNode)

		if listContains(closed, up) {
			up = nilNode
		}
		if listContains(closed, down) {
			down = nilNode
		}
		if listContains(closed, left) {
			left = nilNode
		}
		if listContains(closed, right) {
			right = nilNode
		}

		addOpen := func(n node) {
			if n == nilNode {
				return
			}
			if util.IndexPoint(&track, n.pos) != "#" {
				if listContains(open, n) {
					idx := listIndex(open, n)
					if open[idx].score > n.score {
						open[idx] = n
					}
				} else {
					open = append(open, n)
				}
			}
		}

		addOpen(up)
		addOpen(down)
		addOpen(left)
		addOpen(right)

	}

	panic("no solution")

}

func Day16() {
	track := util.FileToMatrix("input.txt", "", util.NoOpConverter)
	start := util.Find(track, "S")
	end := util.Find(track, "E")

	score := aStar(track, start, end)
	fmt.Println(score)
}
