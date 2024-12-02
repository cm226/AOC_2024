package main

import (
	"bufio"
	"os"
	"strings"
)

func fileToMatrix(fileName string) [][]string {

	inputFile, error := os.Open(fileName)

	if error != nil {
		panic(error)
	}

	scanner := bufio.NewScanner(inputFile)
	matrix := [][]string{}

	for scanner.Scan() {
		line := scanner.Text()
		parts := []string{}

		for _, part := range strings.Split(line, " ") {
			part = strings.TrimSpace(part)
			if part != "" {
				parts = append(parts, part)
			}
		}
		matrix = append(matrix, parts)
	}
	return matrix

}

func abs(first int, second int) int {
	dif := first - second
	if dif < 0 {
		dif *= -1
	}
	return dif
}

func panicIfError(e error) {
	if e != nil {
		panic(e)
	}
}
