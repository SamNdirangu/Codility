package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(Solution([]int{3, 4, 3, 2, 3, -1, 3, 3}))
}

func Solution(someArray []int) int {
	threshold := len(someArray) / 2
	sort.Ints(someArray)

	currentNum := someArray[0]
	currentNumCount := 0

	for index := 0; index < len(someArray); index++ {
		if someArray[index] != currentNum {
			currentNum = someArray[index]
			currentNumCount = 0
		}
		currentNumCount++
		if currentNumCount > threshold {
			break
		}
	}
	return currentNum
}
func solution2(someArray []int) int {
	threshold := len(someArray) / 2
	negativeTrackers := make([]int, len(someArray))
	positiveTrackers := make([]int, len(someArray))
	var denominator int
	for index := 0; index < len(someArray); index++ {
		currentNum := someArray[index]

		if currentNum >= 0 {
			positiveTrackers[currentNum]++
			if positiveTrackers[currentNum] > threshold {
				denominator = currentNum
				break
			}
		} else {
			negativeTrackers[0-currentNum]++
			if negativeTrackers[0-currentNum] > threshold {
				denominator = currentNum
				break
			}
		}
	}
	return denominator
}
