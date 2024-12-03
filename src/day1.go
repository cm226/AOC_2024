package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func pt1(list1 []int, list2 []int) {
	difference_sum := 0
	for i := range len(list1) {
		diff := list1[i] - list2[i]
		if diff < 0 {
			diff *= -1
		}
		difference_sum += diff
	}
	fmt.Println(difference_sum)
}

func pt2(list1Sorted []int, list2Sorted []int) {
	simmilarity := 0

	for _, n := range list1Sorted {
		for _, n2 := range list2Sorted {
			if n2 == n {
				simmilarity += n
			}
		}
	}

	fmt.Print(simmilarity)
}

func day1() {
	list1 := []int{}
	list2 := []int{}

	convertAndAppend := func(list []int, num string) []int {
		first, e := strconv.Atoi(num)
		if e != nil {
			panic(e)
		}
		return append(list, first)
	}

	inputFile, error := os.Open("day1.txt")

	if error != nil {
		panic(error)
	}

	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {
		line := scanner.Text()
		parts := []string{}

		for _, part := range strings.Split(line, " ") {
			part = strings.TrimSpace(part)
			if part != "" {
				parts = append(parts, part)
			}
		}

		list1 = convertAndAppend(list1, parts[0])
		list2 = convertAndAppend(list2, parts[1])
	}

	sort.Ints(list1)
	sort.Ints(list2)

	if len(list1) != len(list2) {
		panic("Lists different lengths")
	}

	pt1(list1, list2)
	pt2(list1, list2)
}
