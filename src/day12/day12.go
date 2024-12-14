package day12

import (
	util "aox_2024/src/utils"
	"fmt"
	"slices"
)

func calcParimeterWithDiscount(region []util.Point, garden [][]string) int {

	// Derp, num continuious lines == number of corners :(, would have been much easier to implement that

	fenceSpace := make([][]int, (len(garden)*2)+1)
	for j := range len(fenceSpace) {
		fenceSpace[j] = make([]int, (len(garden)*2)+1)
	}

	max := util.Point{Y: len(garden), X: len(garden[0])}
	isEdge := func(pt util.Point, veg string) bool {
		if !pt.Inside(max) {
			return true
		}

		if util.IndexPoint(&garden, pt) != veg {
			return true
		}
		return false
	}

	for _, start := range region {
		veg := util.IndexPoint(&garden, start)
		down := start.Add(util.Point{Y: 1})
		up := start.Add(util.Point{Y: -1})
		left := start.Add(util.Point{X: -1})
		right := start.Add(util.Point{X: 1})

		if isEdge(up, veg) {
			fenceSpace[start.Y*2][(start.X*2)+1] = 1
		}

		if isEdge(down, veg) {
			fenceSpace[(start.Y*2)+2][(start.X*2)+1] = 2
		}

		if isEdge(left, veg) {
			fenceSpace[(start.Y*2)+1][start.X*2] = 3
		}

		if isEdge(right, veg) {
			fenceSpace[(start.Y*2)+1][(start.X*2)+2] = 4
		}
	}

	fenceTypes := []int{1, 2, 3, 4}
	discounterPerimeter := 0
	for y := 0; y < len(fenceSpace); y += 2 {
		// count the horizontals
		consecutive := -1
		for x := 1; x < len(fenceSpace[0]); x += 2 {
			if !slices.Contains(fenceTypes, fenceSpace[y][x]) {
				consecutive = -1
				continue
			}

			if consecutive == fenceSpace[y][x] {
				continue
			}

			// is fence and its new
			consecutive = fenceSpace[y][x]
			discounterPerimeter += 1
		}
	}

	for x := 0; x < len(fenceSpace[0]); x += 2 {
		consecutive := -1
		for y := 1; y < len(fenceSpace); y += 2 {
			if !slices.Contains(fenceTypes, fenceSpace[y][x]) {
				consecutive = -1
				continue
			}

			if consecutive == fenceSpace[y][x] {
				continue
			}

			consecutive = fenceSpace[y][x]
			discounterPerimeter += 1
		}
	}

	return discounterPerimeter
}

func calcParimeter(region []util.Point, garden [][]string) int {

	totalPerimeter := 0

	max := util.Point{Y: len(garden), X: len(garden[0])}
	checkPerim := func(pt util.Point, veg string) int {
		if pt.Inside(max) && util.IndexPoint(&garden, pt) == veg {
			return 1
		}
		return 0
	}

	for _, start := range region {
		curPerim := 4

		veg := util.IndexPoint(&garden, start)
		up := start.Add(util.Point{Y: 1})
		down := start.Add(util.Point{Y: -1})
		left := start.Add(util.Point{X: -1})
		right := start.Add(util.Point{X: 1})

		curPerim -= checkPerim(up, veg)
		curPerim -= checkPerim(down, veg)
		curPerim -= checkPerim(left, veg)
		curPerim -= checkPerim(right, veg)
		totalPerimeter += curPerim
	}
	return totalPerimeter
}

func findRegion(start util.Point, region *[]util.Point, garden [][]string) {

	veg := util.IndexPoint(&garden, start)
	up := start.Add(util.Point{Y: 1})
	down := start.Add(util.Point{Y: -1})
	left := start.Add(util.Point{X: -1})
	right := start.Add(util.Point{X: 1})

	max := util.Point{Y: len(garden), X: len(garden[0])}
	checkVeg := func(pt util.Point) {
		if pt.Inside(max) && util.IndexPoint(&garden, pt) == veg && !slices.Contains(*region, pt) {
			*region = append(*region, pt)
			findRegion(pt, region, garden)
		}
	}
	checkVeg(start)
	checkVeg(up)
	checkVeg(down)
	checkVeg(left)
	checkVeg(right)

}

func getRegions(garden [][]string) [][]util.Point {

	regions := [][]util.Point{}
	nextRegionStart := util.Point{X: 0, Y: 0}
	nilPoint := util.Point{X: -1, Y: -1}
	for nextRegionStart != nilPoint {
		newRegion := []util.Point{}
		findRegion(nextRegionStart, &newRegion, garden)

		regions = append(regions, newRegion)

		lastRegion := nextRegionStart
		nextRegionStart = nilPoint
		for i := range garden {
			for j := range garden {
				if (i*len(garden) + j) < (lastRegion.Y*len(garden))+lastRegion.X {
					continue
				}

				found := false
				for _, region := range regions {
					if slices.Contains(region, util.Point{Y: i, X: j}) {
						found = true
						break
					}
				}
				if !found {
					nextRegionStart = util.Point{X: j, Y: i}
					break
				}
			}
			if nextRegionStart != nilPoint {
				break
			}
		}
	}

	return regions

}

func Day12() {
	garden := util.FileToMatrix("input.txt", "", util.NoOpConverter)

	regions := getRegions(garden)

	//part1
	totalCost := 0
	for _, region := range regions {
		totalCost += calcParimeter(region, garden) * len(region)
	}

	fmt.Println(totalCost)

	totalCost = 0
	for _, region := range regions {
		totalCost += calcParimeterWithDiscount(region, garden) * len(region)
	}

	fmt.Println(totalCost)
}
