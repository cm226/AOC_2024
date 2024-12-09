package day9

import (
	util "aox_2024/src/utils"
	"fmt"
	"strconv"
)

func compactPt1(expandedDisk []int) {

	firstSpace := 0
	lastFile := len(expandedDisk) - 1

	for firstSpace < lastFile {
		for i := firstSpace; i < len(expandedDisk); i++ {
			if expandedDisk[i] == -1 {
				firstSpace = i
				break
			}
		}

		for i := lastFile; i > 0; i-- {
			if expandedDisk[i] != -1 {
				lastFile = i
				break
			}
		}

		if firstSpace >= lastFile {
			break
			// done
		}

		expandedDisk[firstSpace] = expandedDisk[lastFile]
		expandedDisk[lastFile] = -1
	}

	// checksum
	checksum := 0
	for i, b := range expandedDisk {
		if b != -1 {
			checksum += i * b
		}
	}
	fmt.Println(checksum)
}

func compactPt2(du []int, expandedDisk []int) {

	lastFile := len(expandedDisk) - 1
	lastID := int(^uint(0) >> 1)

	for lastFile > 0 {
		for i := lastFile; i >= 0; i-- {
			if expandedDisk[i] < lastID && expandedDisk[i] != -1 {
				lastFile = i
				lastID = expandedDisk[i]
				break
			}
			if i == 0 {
				lastFile = 0
			}
		}
		if lastFile == 0 {
			break //done
		}

		sizeOfFile := du[lastID*2]
		contigious := 0
		for i := range lastFile {
			if expandedDisk[i] == -1 {
				contigious += 1
			} else {
				contigious = 0
			}
			if contigious >= sizeOfFile {
				for s := range sizeOfFile {
					expandedDisk[i-s] = expandedDisk[lastFile-s]
					expandedDisk[lastFile-s] = -1
				}
				break
			}
		}
	}

	// checksum
	checksum := 0
	for i, b := range expandedDisk {
		if b != -1 {
			checksum += i * b
		}
	}
	fmt.Println(checksum)
}

func part1(du []int) {
	expandedDisk := []int{}
	fileID := 0

	for i := 0; i < len(du); i += 2 {
		blockSize := du[i]
		freeSpace := 0
		if i <= len(du)-2 {
			freeSpace = du[i+1]
		}

		for range blockSize {
			expandedDisk = append(expandedDisk, fileID)
		}

		for range freeSpace {
			expandedDisk = append(expandedDisk, -1)
		}
		fileID += 1
	}

	//compactPt1(expandedDisk)
	compactPt2(du, expandedDisk)

}

func Day9() {
	du := util.FileToSlice("input.txt", func(s string) int {
		i, e := strconv.Atoi(s)
		util.PanicIfError(e)
		return i
	})
	part1(du)
}
