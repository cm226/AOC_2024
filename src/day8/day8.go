package day8

import (
	util "aox_2024/src/utils"
	"fmt"
)

func antinodespt2(antenaPairs map[string][]util.Point, inputSize int) {

	uniquePoints := make([][]int, inputSize)

	for i := range inputSize {
		uniquePoints[i] = make([]int, inputSize)
	}

	for _, antena := range antenaPairs {
		for i := range len(antena) {
			for j := i + 1; j < len(antena); j++ {
				antinode1Delta := antena[i].Sub(antena[j])
				antinodeLocation1 := antena[i].Add(antinode1Delta)

				uniquePoints[antena[i].Y][antena[i].X] = 1
				uniquePoints[antena[j].Y][antena[j].X] = 1

				for antinodeLocation1.Inside(util.Point{X: inputSize, Y: inputSize}) {
					uniquePoints[antinodeLocation1.Y][antinodeLocation1.X] = 1

					antinodeLocation1 = antinodeLocation1.Add(antinode1Delta)
				}
				antinode2Delta := antena[j].Sub(antena[i])
				antinodeLocation2 := antena[j].Add(antinode2Delta)

				for antinodeLocation2.Inside(util.Point{X: inputSize, Y: inputSize}) {
					uniquePoints[antinodeLocation2.Y][antinodeLocation2.X] = 1
					antinodeLocation2 = antinodeLocation2.Add(antinode2Delta)
				}
			}
		}
	}

	total := 0
	for _, line := range uniquePoints {
		fmt.Println(line)
		for _, i := range line {
			total += i
		}
	}
	fmt.Println(total)
}

func antinodespt1(antenaPairs map[string][]util.Point, inputSize int) {

	antinodeMap := map[string][]util.Point{}
	uniquePoints := make([][]int, inputSize)

	for i := range inputSize {
		uniquePoints[i] = make([]int, inputSize)
	}

	for aname, antena := range antenaPairs {
		for i := range len(antena) {
			for j := i + 1; j < len(antena); j++ {
				// get delta
				antinodeLocation1 := util.Point{
					X: antena[i].X + (antena[i].X - antena[j].X),
					Y: antena[i].Y + (antena[i].Y - antena[j].Y),
				}

				antinodeLocation2 := util.Point{
					X: antena[j].X + (antena[j].X - antena[i].X),
					Y: antena[j].Y + (antena[j].Y - antena[i].Y),
				}

				if antinodeLocation1.X >= 0 && antinodeLocation1.X < inputSize && antinodeLocation1.Y >= 0 && antinodeLocation1.Y < inputSize {
					antinodeMap[aname] = append(antinodeMap[aname], antinodeLocation1)
					uniquePoints[antinodeLocation1.Y][antinodeLocation1.X] = 1
				}

				if antinodeLocation2.X >= 0 && antinodeLocation2.X < inputSize && antinodeLocation2.Y >= 0 && antinodeLocation2.Y < inputSize {
					antinodeMap[aname] = append(antinodeMap[aname], antinodeLocation2)
					uniquePoints[antinodeLocation2.Y][antinodeLocation2.X] = 1
				}
			}
		}
	}

	total := 0
	for _, line := range uniquePoints {
		fmt.Println(line)
		for _, i := range line {
			total += i
		}
	}
	fmt.Println(total)
}

func antenaPairs(antenaMap [][]string) {

	antenaTypes := map[string][]util.Point{}

	for i, line := range antenaMap {
		for j, c := range line {
			if c != "." {
				antenaTypes[c] = append(antenaTypes[c], util.Point{X: j, Y: i})
			}
		}
	}
	antinodespt2(antenaTypes, len(antenaMap))
}

func part1(antenaMap [][]string) {
	antenaPairs(antenaMap)
}

func Day8() {
	antenaMap := util.FileToMatrix("input.txt", "", util.NoOpConverter)
	part1(antenaMap)

}
