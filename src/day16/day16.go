package day16

import (
	util "aox_2024/src/utils"
	"fmt"
	"slices"
)

const MaxUint = ^uint(0)
const MaxInt = int(MaxUint >> 1)

type node struct {
	pos       util.Point
	direction int
	score     int
	previous  []node
}

func listContains(list *[]node, n node) bool {

	sFn := func(snode node) bool {
		if snode.pos == n.pos && snode.direction == n.direction {
			return true
		}
		return false
	}
	if slices.ContainsFunc(*list, sFn) {
		return true
	}
	return false
}

func listIndex(list []node, n node) int {

	iFn := func(sNode node) bool {
		if sNode.pos == n.pos && sNode.direction == n.direction {
			return true
		}
		return false
	}
	return slices.IndexFunc(list, iFn)
}

func comparePrev(p1 []node, p2 []node) bool {
	if len(p1) != len(p2) {
		return false
	}

	for i := range p1 {
		if p1[i].pos != p2[i].pos {
			return false
		}
	}
	return true
}

func aStar(track [][]string, start util.Point, end util.Point) []node {

	open := []node{
		node{
			pos:       start,
			direction: 1,
			score:     0,
		},
	}
	closed := []node{}

	paths := []node{}

	lowestScore := MaxInt

	for len(open) != 0 {
		lowest := open[0].score
		currentNode := open[0]
		openi := 0
		for i, n := range open {
			if n.score < lowest {
				currentNode = n
				openi = i
				lowest = n.score
			}
		}

		open = slices.Delete(open, openi, openi+1)
		closed = append(closed, currentNode)

		if currentNode.score > lowestScore {
			continue
		}

		if currentNode.pos == end {
			paths = append(paths, currentNode)
			if currentNode.score < lowestScore {
				lowestScore = currentNode.score
			}
			continue
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
			previous:  append(slices.Clone(currentNode.previous), currentNode),
		}
		down := node{
			pos:       currentNode.pos.Add(util.PDown()),
			direction: 2,
			previous:  append(slices.Clone(currentNode.previous), currentNode),
		}
		right := node{
			pos:       currentNode.pos.Add(util.PRight()),
			direction: 1,
			previous:  append(slices.Clone(currentNode.previous), currentNode),
		}
		left := node{
			pos:       currentNode.pos.Add(util.PLeft()),
			direction: 3,
			previous:  append(slices.Clone(currentNode.previous), currentNode),
		}
		up.score = calcScore(up, currentNode)
		down.score = calcScore(down, currentNode)
		left.score = calcScore(left, currentNode)
		right.score = calcScore(right, currentNode)

		if listContains(&closed, up) {
			up = nilNode
		}
		if listContains(&closed, down) {
			down = nilNode
		}
		if listContains(&closed, left) {
			left = nilNode
		}
		if listContains(&closed, right) {
			right = nilNode
		}

		addOpen := func(n node) {
			if n.score == nilNode.score {
				return
			}
			if util.IndexPoint(&track, n.pos) != "#" {

				if listContains(&open, n) {
					idx := listIndex(open, n)

					if comparePrev(open[idx].previous, n.previous) {

						if open[idx].score > n.score {
							open[idx] = n
						}
					} else {
						open = append(open, n)
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

	return paths

}

func Day16() {
	track := util.FileToMatrix("input.txt", "", util.NoOpConverter)
	start := util.Find(track, "S")
	end := util.Find(track, "E")

	paths := aStar(track, start, end)

	grid := make([][]int, len(track))
	for i := range grid {
		grid[i] = make([]int, len(track[0]))
	}

	for _, path := range paths {
		for _, p := range path.previous {
			grid[p.pos.Y][p.pos.X] = 1
		}
	}
	grid[end.Y][end.X] = 1

	total := 0
	for i := range grid {
		for j := range grid[i] {
			total += grid[i][j]
		}
	}

	// pt 1
	fmt.Println(paths[len(paths)-1].score)
	// pt 2
	fmt.Println(total)
}
