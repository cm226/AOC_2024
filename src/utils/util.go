package util

import (
	"bufio"
	"os"
	"strings"
)

func NoOpConverter(s string) string {
	return s
}

func FileToSlice(fileName string) []string {
	inputFile, error := os.Open(fileName)

	if error != nil {
		panic(error)
	}

	scanner := bufio.NewScanner(inputFile)
	chars := []string{}

	for scanner.Scan() {
		line := scanner.Text()
		chars = append(chars, strings.Split(line, "")...)
	}
	return chars
}

func FileToMatrix[T any](fileName string, sep string, converter func(string) T) [][]T {

	inputFile, error := os.Open(fileName)

	if error != nil {
		panic(error)
	}

	scanner := bufio.NewScanner(inputFile)
	matrix := [][]T{}

	for scanner.Scan() {
		line := scanner.Text()
		parts := []T{}

		for _, part := range strings.Split(line, sep) {
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

func Abs(first int, second int) int {
	dif := first - second
	if dif < 0 {
		dif *= -1
	}
	return dif
}

func PanicIfError(e error) {
	if e != nil {
		panic(e)
	}
}
