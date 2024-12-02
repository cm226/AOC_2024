package main

import (
	"bufio"
	"os"
	"strings"
)

func noOpConverter(s string) string {
	return s
}

func fileToMatrix[T any](fileName string, converter func(string) T) [][]T {

	inputFile, error := os.Open(fileName)

	if error != nil {
		panic(error)
	}

	scanner := bufio.NewScanner(inputFile)
	matrix := [][]T{}

	for scanner.Scan() {
		line := scanner.Text()
		parts := []T{}

		for _, part := range strings.Split(line, " ") {
			part = strings.TrimSpace(part)
			convertedPart := converter(part)
			if part != "" {
				parts = append(parts, convertedPart)
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
