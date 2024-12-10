package util

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Point struct {
	X int
	Y int
}

func (p Point) Add(p2 Point) Point {
	return Point{
		X: p.X + p2.X,
		Y: p.Y + p2.Y,
	}
}

func (p Point) Sub(p2 Point) Point {
	return Point{
		X: p.X - p2.X,
		Y: p.Y - p2.Y,
	}
}

func (p Point) Inside(max Point) bool {
	return p.X < max.X && p.Y < max.Y && p.X >= 0 && p.Y >= 0
}

func IndexPoint[T any](matrix *[][]T, point Point) T {
	return (*matrix)[point.Y][point.X]
}

func NoOpConverter(s string) string {
	return s
}

func FileToSlice[T any](fileName string, converter func(string) T) []T {
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

	ret := make([]T, len(chars))
	for i, c := range chars {
		ret[i] = converter(c)
	}
	return ret
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

func PrintMatrix[T any](matrix [][]T) {
	for i := range len(matrix) {
		fmt.Println(matrix[i])
	}
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
