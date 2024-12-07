package day7

import (
	util "aox_2024/src/utils"
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type eq struct {
	result int
	nums   []int
}

type ops int

const (
	Add ops = iota
	Mult
	Concat
)

func parseInput(fileName string) []eq {

	inputFile, error := os.Open(fileName)

	if error != nil {
		panic(error)
	}

	scanner := bufio.NewScanner(inputFile)
	eqs := []eq{}

	for scanner.Scan() {
		line := scanner.Text()
		res_num := strings.Split(line, ":")

		if len(res_num) != 2 {
			panic("failed to parse input")
		}

		nums_str := strings.Split(strings.TrimSpace(res_num[1]), " ")
		nums := []int{}
		for _, num_str := range nums_str {
			num_i, e := strconv.Atoi(num_str)
			util.PanicIfError(e)
			nums = append(nums, num_i)
		}

		result, e := strconv.Atoi(res_num[0])
		util.PanicIfError(e)

		eqs = append(eqs, eq{
			result,
			nums,
		})
	}
	return eqs
}

func checkPasses(eq eq, ops []ops) bool {
	total := eq.nums[0]
	for i := range len(eq.nums) - 1 {
		op := ops[i]

		switch op {
		case Add:
			total = total + eq.nums[i+1]
		case Mult:
			total = total * eq.nums[i+1]
		case Concat:
			numStr := strconv.FormatInt(int64(total), 10) + strconv.FormatInt(int64(eq.nums[i+1]), 10)
			total2, e := strconv.Atoi(numStr)
			util.PanicIfError(e)
			total = total2
		}
	}

	return total == eq.result
}

func genPerms(eq eq) []ops {
	opArr := make([]ops, len(eq.nums)-1)
	count := 0
	numPerms := int(math.Pow(4, float64(len(opArr))))
	for range numPerms {

		// convert count to ops
		count_copy := count
		opArr = make([]ops, len(eq.nums)-1)
		skip := false
		for j := range len(opArr) {
			if count_copy&1 != 0 {
				opArr[j] = Mult
			} else if count_copy&2 != 0 {
				opArr[j] = Add
			} else if count_copy&3 == 0 {
				opArr[j] = Concat
			} else if count_copy&3 == 3 {
				// skip this one
				skip = true
				break
			}
			count_copy = count_copy >> 2
		}
		if skip {
			continue
		}
		if checkPasses(eq, opArr) {
			return opArr
		}

		count += 1
	}
	return nil
}

func part2(eqs []eq) {
	total := 0
	for _, eq := range eqs {
		ops := genPerms(eq)
		if ops != nil {
			total += eq.result
		}
	}
	fmt.Println(total)

}

func Day7() {
	eqs := parseInput("input.txt")
	part2(eqs)
}
