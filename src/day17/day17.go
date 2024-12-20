package day17

import (
	util "aox_2024/src/utils"
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

const MaxUint = ^uint(0)
const MinUint = 0
const MaxInt = int(MaxUint >> 1)
const MinInt = -MaxInt - 1

func parse(fileName string) (int, int, int, []int) {
	input, e := os.Open(fileName)

	util.PanicIfError(e)

	scanner := bufio.NewScanner(input)

	scanner.Scan()
	regA := scanner.Text()[len("Register A: "):]
	scanner.Scan()
	regB := scanner.Text()[len("Register A: "):]
	scanner.Scan()
	regC := scanner.Text()[len("Register A: "):]

	scanner.Scan() // skip empty

	scanner.Scan()
	programStr := scanner.Text()[len("Program: "):]

	regAVal, e := strconv.Atoi(regA)
	util.PanicIfError(e)

	regBVal, e := strconv.Atoi(regB)
	util.PanicIfError(e)

	regCVal, e := strconv.Atoi(regC)
	util.PanicIfError(e)

	programStrs := strings.Split(programStr, ",")
	program := []int{}
	for ps := range programStrs {
		i, e := strconv.Atoi(programStrs[ps])
		util.PanicIfError(e)
		program = append(program, i)
	}

	return regAVal, regBVal, regCVal, program
}

func valFromOp(reg *[]int, opp int) int {
	if opp > 6 {
		panic("invalid op")
	}
	if opp <= 3 {
		return opp
	}
	return (*reg)[opp-4]

}

func adv(reg *[]int, program []int, iPtr *int) {

	operand := program[*iPtr]
	*iPtr += 1

	pow := math.Pow(2.0, float64(valFromOp(reg, operand)))

	res := float64((*reg)[0]) / pow

	(*reg)[0] = int(math.Floor(res))
}

func bxl(reg *[]int, program []int, iPtr *int) {

	operand := program[*iPtr]
	*iPtr += 1
	res := (*reg)[1] ^ operand
	(*reg)[1] = res
}

func bst(reg *[]int, program []int, iPtr *int) {

	operand := valFromOp(reg, program[*iPtr])
	*iPtr += 1

	res := operand % 8
	(*reg)[1] = res
}

func jnz(reg *[]int, program []int, iPtr *int) {
	if (*reg)[0] == 0 {
		*iPtr++
		return
	}

	*iPtr = program[*iPtr]
}

func bxc(reg *[]int, program []int, iPtr *int) {

	*iPtr += 1
	res := (*reg)[1] ^ (*reg)[2]
	(*reg)[1] = res
}

func out(reg *[]int, program []int, iPtr *int, outputs *[]int) {

	operand := program[*iPtr]
	*iPtr += 1
	res := valFromOp(reg, operand) % 8
	*outputs = append(*outputs, res)
}

func bdv(reg *[]int, program []int, iPtr *int) {

	operand := program[*iPtr]
	*iPtr += 1
	res := float64((*reg)[0]) / math.Pow(2.0, float64(valFromOp(reg, operand)))
	(*reg)[1] = int(math.Floor(res))
}

func cdv(reg *[]int, program []int, iPtr *int) {

	operand := program[*iPtr]
	*iPtr += 1

	res := float64((*reg)[0]) / math.Pow(2.0, float64(valFromOp(reg, operand)))
	(*reg)[2] = int(math.Floor(res))
}
func runProgram(registers []int, program []int, targetOutput []int) []int {

	outputs := []int{}
	iPtr := 0

	for iPtr < len(program) {

		opcode := program[iPtr]
		iPtr += 1

		switch opcode {
		case 0:
			adv(&registers, program, &iPtr)
		case 1:
			bxl(&registers, program, &iPtr)
		case 2:
			bst(&registers, program, &iPtr)
		case 3:
			jnz(&registers, program, &iPtr)
		case 4:
			bxc(&registers, program, &iPtr)
		case 5:
			out(&registers, program, &iPtr, &outputs)
			if len(targetOutput) != 0 { // pt 2
				for i := range outputs {
					if targetOutput[i] != outputs[i] {
						return []int{}
					}
				}
			}
		case 6:
			bdv(&registers, program, &iPtr)
		case 7:
			cdv(&registers, program, &iPtr)
		}
	}
	return outputs
}

func Day17() {
	regA, regB, regC, program := parse("input.txt")

	registers := []int{regA, regB, regC}

	//outputs := runProgram(registers, program, []int{})

	//for _, o := range outputs {
	//fmt.Print(o)
	//fmt.Print(",")
	//}
	//fmt.Println()
	// part 2
	//aReg := MaxInt
	//for !reflect.DeepEqual(outputs, program) {
	//registers[0] = aReg
	//outputs = runProgram(registers, program, program)

	//if aReg%10000 == 0 {
	//fmt.Println(aReg)
	//}

	//aReg--
	//testI := aReg >> 3
	//for testI%8 != 2 {
	//aReg--
	//testI = aReg >> 3
	//if aReg == 0 {
	//panic("Failed, missed it")
	//}
	//}
	//}

	// target is 2,4,1,1,7,5,4,4,1,4,0,3,5,5,3,0
	// we need A >> 3 last 3 bits == 2
	// then A >> 6 last 3 bits == 4

	// Sda face I mis read the program it actually prints b :( and now i dont want to re-do, soooo ill do it later :D maybe

	// so working in reverse

	// A = 010 // 2
	// A = 100 010 // 4
	// etc

	A := 1
	rProgram := slices.Clone(program)
	slices.Reverse(rProgram)
	for _, target := range rProgram {
		A = A << 3
		A = A | target
	}
	A = A << 3
	fmt.Println(A)

	registers[0] = 130502792549136
	outputs := runProgram(registers, program, []int{})

	for _, o := range outputs {
		fmt.Print(o)
		fmt.Print(",")
	}
	fmt.Println()
}
