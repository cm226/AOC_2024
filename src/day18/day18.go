package day18

import (
	util "aox_2024/src/utils"
	"fmt"
	"slices"
	"strconv"
)

type node struct {
	pos           util.Point
	count         int
	scoreEsitmate int
}

func listIndex(nodes *[]node, n node) int {
	for i, cn := range *nodes {
		if cn.pos == n.pos {
			return i
		}
	}
	panic("not fount")
}

func listContains(nodes *[]node, n node) bool {
	for _, curNode := range *nodes {
		if curNode.pos == n.pos {
			return true
		}
	}
	return false
}

func listContainsPoint(points []util.Point, p util.Point) bool {
	for _, point := range points {
		if point == p {
			return true
		}
	}
	return false
}
func aStar(corrupted []util.Point, start util.Point, end util.Point) node {

	open := []node{
		node{
			pos:           start,
			count:         0,
			scoreEsitmate: ((end.X - start.X) << 2) + ((end.Y - start.Y) << 2),
		},
	}
	closed := []node{}

	for len(open) != 0 {
		lowest := open[0].scoreEsitmate
		currentNode := open[0]
		openi := 0
		for i, n := range open {
			if n.count < lowest {
				currentNode = n
				openi = i
				lowest = n.count
			}
		}

		open = slices.Delete(open, openi, openi+1)
		closed = append(closed, currentNode)

		if currentNode.pos == end {
			return currentNode
		}

		// adjacents
		up := node{
			pos:   currentNode.pos.Add(util.PUp()),
			count: currentNode.count + 1,
		}
		up.scoreEsitmate = ((end.X - up.pos.X) << 2) + ((end.Y - up.pos.Y) << 2)

		down := node{
			pos:   currentNode.pos.Add(util.PDown()),
			count: currentNode.count + 1,
		}
		down.scoreEsitmate = ((end.X - down.pos.X) << 2) + ((end.Y - down.pos.Y) << 2)

		right := node{
			pos:   currentNode.pos.Add(util.PRight()),
			count: currentNode.count + 1,
		}
		right.scoreEsitmate = ((end.X - right.pos.X) << 2) + ((end.Y - right.pos.Y) << 2)

		left := node{
			pos:   currentNode.pos.Add(util.PLeft()),
			count: currentNode.count + 1,
		}
		left.scoreEsitmate = ((end.X - left.pos.X) << 2) + ((end.Y - left.pos.Y) << 2)

		addOpen := func(n node) {
			if listContains(&closed, n) {
				return
			}
			if listContainsPoint(corrupted, n.pos) == false && n.pos.Inside(util.Point{X: 71, Y: 71}) {

				if listContains(&open, n) {
					idx := listIndex(&open, n)

					if open[idx].count > n.count {
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

	return node{
		count: -1,
	}
}

func Day18() {
	input := util.FileToMatrix("input.txt", ",", func(s string) int {
		i, e := strconv.Atoi(s)
		util.PanicIfError(e)
		return i
	})

	memory := []util.Point{}
	for _, line := range input {
		memory = append(memory, util.Point{
			X: line[0],
			Y: line[1],
		})
	}

	// pt 1
	endNode := aStar(memory[0:1024], util.Point{X: 0, Y: 0}, util.Point{X: 70, Y: 70})
	fmt.Println(endNode.count)

	// pt2
	for i := 1024; i < len(memory); i++ {
		result := aStar(memory[0:i], util.Point{X: 0, Y: 0}, util.Point{X: 70, Y: 70})
		if result.count == -1 {
			fmt.Println(memory[i-1])
			break

		}
	}

}
