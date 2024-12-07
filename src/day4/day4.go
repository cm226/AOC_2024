package day4

import (
	util "aox_2024/src/utils"
	"fmt"
	"iter"
	"slices"
)

func findInLine(line iter.Seq2[int, string]) int {
	// yuck :( please nobody look at this :D

	searchTerm := []string{"M", "A", "S"}
	tmp := make([]string, len(searchTerm))
	copy(tmp, searchTerm)

	searchCount := 0
	for _, c := range line {
		if c == tmp[0] {
			tmp = tmp[1:]
		} else {
			//reset
			if len(tmp) != len(searchTerm) {
				tmp = make([]string, len(searchTerm))
				copy(tmp, searchTerm)
				if c == tmp[0] {
					tmp = tmp[1:]
				}
			}
		}
		if len(tmp) == 0 {
			// found 1
			searchCount += 1
			tmp = make([]string, len(searchTerm))
			copy(tmp, searchTerm)
		}
	}

	return searchCount
}

func makeDiag(puzzle [][]string, start int, dir int) []string {
	line := []string{}
	for i := range len(puzzle) {
		if start < len(puzzle[i]) && start >= 0 {
			line = append(line, puzzle[i][start])
		}
		start += dir
	}
	return line
}

func transpose(puzzle [][]string) [][]string {
	transposed := make([][]string, len(puzzle))
	for i := range len(puzzle) {
		transposed[i] = make([]string, len(puzzle[i]))
	}

	for i := range len(puzzle) {
		for j := range puzzle[i] {
			transposed[j][i] = puzzle[i][j]
		}
	}
	return transposed
}

func day4Part1(puzzle [][]string) {

	totalCount := 0
	for _, line := range puzzle {
		totalCount += findInLine(slices.All(line))
		totalCount += findInLine(slices.Backward(line))
	}

	for i := range len(puzzle[0]) * 2 {
		totalCount += findInLine(slices.All(makeDiag(puzzle, i, -1)))
		totalCount += findInLine(slices.Backward(makeDiag(puzzle, i, -1)))
	}

	for i := range len(puzzle[0]) * 2 {
		j := i - len(puzzle[0])
		totalCount += findInLine(slices.All(makeDiag(puzzle, j, 1)))
		totalCount += findInLine(slices.Backward(makeDiag(puzzle, j, 1)))
	}

	puzzleT := transpose(puzzle)

	for _, line := range puzzleT {
		totalCount += findInLine(slices.All(line))
		totalCount += findInLine(slices.Backward(line))
	}
	fmt.Println(totalCount)

}

func day4Part2(puzzle [][]string) {
	totalCount := 0

	for i, line := range puzzle {
		for j, c := range line {

			if c == "A" {
				if i >= 1 && i <= len(puzzle)-2 && j >= 1 && j <= len(puzzle[i])-2 {

					diag1 := []string{puzzle[i-1][j-1], c, puzzle[i+1][j+1]}
					diag2 := []string{puzzle[i-1][j+1], c, puzzle[i+1][j-1]}

					diag1count := findInLine(slices.All(diag1))
					diag1count += findInLine(slices.Backward(diag1))

					diag2Count := findInLine(slices.All(diag2))
					diag2Count += findInLine(slices.Backward(diag2))

					if diag1count > 0 && diag2Count > 0 {
						totalCount += 1
					}
				}
			}
		}
	}
	fmt.Println(totalCount)

}

func Day4() {
	puzzle := util.FileToMatrix("input.txt", "", util.NoOpConverter)
	day4Part1(puzzle)
	day4Part2(puzzle)
}
