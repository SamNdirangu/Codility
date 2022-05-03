package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	fmt.Println(Solution([]int{2}))
}

//Easiest if returning the denominator not index
func Solution(someArray []int) int {
	sortedArray := someArray
	sort.Ints(sortedArray)

	//variables
	denominator := -1
	searchDenom := false

	currentNum := 0
	currentNumCount := 0
	threshold := math.Ceil(len(someArray) / 2)

	for index := 0; index < len(sortedArray); index++ {
		if sortedArray[index] != currentNum {
			currentNum = sortedArray[index]
			currentNumCount = 0
		}
		currentNumCount++
		if currentNumCount > threshold {
			searchDenom = true
			break
		}
	}
	if !searchDenom {
		return denominator
	}
	for index := 0; index < len(someArray); index++ {
		if someArray[index] == currentNum {
			denominator = index
			break
		}
	}
	return denominator
}

//func solution2(someArray []int) int {
//	threshold := len(someArray) / 2
//	negativeTrackers := make([]int, len(someArray))
//	positiveTrackers := make([]int, len(someArray))
//	var denominator int
//	for index := 0; index < len(someArray); index++ {
//		currentNum := someArray[index]
//
//		if currentNum >= 0 {
//			positiveTrackers[currentNum]++
//			if positiveTrackers[currentNum] > threshold {
//				denominator = currentNum
//				break
//			}
//		} else {
//			negativeTrackers[0-currentNum]++
//			if negativeTrackers[0-currentNum] > threshold {
//				denominator = currentNum
//				break
//			}
//		}
//	}
//	return denominator
//}
