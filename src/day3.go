package main

import (
	util "aox_2024/src/utils"
	"fmt"
	"strconv"
	"strings"
)

func isToken(input []string, sequence []string) bool {
	for i, c := range sequence {
		if input[i] != c {
			return false
		}
	}
	return true
}

func isIntToken(input []string) int {
	i := 1
	_, e := strconv.Atoi(input[i-1])
	if e != nil {
		return 0
	}

	for e == nil {
		_, e = strconv.Atoi(input[i])
		if e == nil {
			i++
		}

	}
	return i
}

func day3_part1_and_part_2(input []string) {

	finalValue := 0

	doToken := []string{"d", "o", "(", ")"}
	dontToken := []string{"d", "o", "n", "'", "t", "(", ")"}
	mulToken := []string{"m", "u", "l", "("}

	enabled := true

	for i := range len(input) {
		if isToken(input[i:], doToken) {
			enabled = true
			continue
		}
		if isToken(input[i:], dontToken) {
			enabled = false
			continue
		}
		if isToken(input[i:], mulToken) && enabled {
			firstArgLen := isIntToken(input[i+len(mulToken):])

			if firstArgLen < 1 || firstArgLen > 3 {
				continue
			}

			if input[i+len(mulToken)+firstArgLen] != "," {
				continue
			}

			secondArgLen := isIntToken(input[i+len(mulToken)+firstArgLen+1:])

			if secondArgLen < 1 || secondArgLen > 3 {
				continue
			}

			if input[i+len(mulToken)+firstArgLen+1+secondArgLen] != ")" {
				continue
			}

			// found 1 parse it
			firstArg := strings.Join(input[i+len(mulToken):i+len(mulToken)+firstArgLen], "")
			secondArg := strings.Join(input[i+len(mulToken)+firstArgLen+1:i+len(mulToken)+firstArgLen+1+secondArgLen], "")

			arg1, e := strconv.Atoi(firstArg)
			util.PanicIfError(e)
			arg2, e := strconv.Atoi(secondArg)
			util.PanicIfError(e)

			finalValue += arg1 * arg2
		}
	}

	fmt.Println(finalValue)

}

func day3() {

	file := util.FileToSlice("input.txt")
	fmt.Println(file)

	day3_part1_and_part_2(file)
}
