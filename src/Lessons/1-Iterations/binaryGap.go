package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(solution(1202))
}

func solution(someNumber int) int {
	numberBinary := strconv.FormatInt(int64(someNumber), 2) //convert to binary
	fmt.Println(numberBinary)
	gapSize := 0
	tempGapSize := 0

	for index := 0; index < len(numberBinary); index++ {
		if string(numberBinary[index]) == "1" {
			if tempGapSize > gapSize {
				gapSize = tempGapSize
			}
			tempGapSize = 0
		} else {
			tempGapSize++
		}
	}
	return gapSize
}
