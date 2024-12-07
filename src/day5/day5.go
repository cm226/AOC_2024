package day5

import (
	util "aox_2024/src/utils"
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type rule struct {
	first  int
	second int
}

func parseInput(fileName string) ([]rule, [][]int) {

	inputFile, error := os.Open(fileName)

	if error != nil {
		panic(error)
	}

	scanner := bufio.NewScanner(inputFile)
	rules := []rule{}
	allpages := [][]int{}

	parsingRules := false
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			parsingRules = true
			continue
		}

		if !parsingRules {
			splitLine := strings.Split(line, "|")
			if len(splitLine) != 2 {
				panic("failed to parse input")
			}

			first, e := strconv.Atoi(splitLine[0])
			util.PanicIfError(e)
			second, e := strconv.Atoi(splitLine[1])
			util.PanicIfError(e)
			rule := rule{
				first:  first,
				second: second,
			}
			rules = append(rules, rule)
		} else {
			pagesStr := strings.Split(line, ",")
			pages := []int{}
			for _, pageStr := range pagesStr {
				page, e := strconv.Atoi(pageStr)
				util.PanicIfError(e)
				pages = append(pages, page)
			}
			allpages = append(allpages, pages)
		}
	}
	return rules, allpages
}

func checkRule(rule rule, page []int) bool {
	if !(slices.Contains(page, rule.first) && slices.Contains(page, rule.second)) {
		return true // rules dosnt apply
	}

	firstI := slices.Index(page, rule.first)
	secondI := slices.Index(page, rule.second)
	return firstI < secondI
}

func checkPageOrder(rules []rule, page []int) bool {

	for _, rule := range rules {
		if !checkRule(rule, page) {
			return false
		}
	}
	return true
}

func fixAPage(rules []rule, page []int) []int {

	newPage := []int{page[0]}
	for _, p := range page[1:] {
		for i := range len(newPage) + 1 {
			tmpNewPage := make([]int, len(newPage))
			copy(tmpNewPage, newPage)
			tmpNewPage = slices.Insert(tmpNewPage, i, p)
			if checkPageOrder(rules, tmpNewPage) {
				newPage = tmpNewPage
				break
			}
		}
	}

	return newPage
}

func day5_part1(rules []rule, pages [][]int) {

	correctlyOrderedSum := 0
	for _, page := range pages {
		if checkPageOrder(rules, page) {
			if len(page)%2 == 0 {
				panic("even page length !?")
			}
			correctlyOrderedSum += page[(len(page)-1)/2]
		}
	}

	fmt.Println(correctlyOrderedSum)
}

func day5_part2(rules []rule, pages [][]int) {

	correctlyOrderedSum := 0
	for _, page := range pages {
		if !checkPageOrder(rules, page) {

			if len(page)%2 == 0 {
				panic("even page length !?")
			}
			newPage := fixAPage(rules, page)
			correctlyOrderedSum += newPage[(len(newPage)-1)/2]
		}
	}

	fmt.Println(correctlyOrderedSum)
}

func Day5() {
	rules, pages := parseInput("input.txt")
	day5_part2(rules, pages)
}
