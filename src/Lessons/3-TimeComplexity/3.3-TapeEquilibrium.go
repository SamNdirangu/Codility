package main

import (
	"fmt"
	"math"
)

func main() {

	fmt.Println(solution([]int{3, 1, 2, 4, 3}))
}

func solution(someArray []int) int {
	rightHandSum := 0 //store our sums
	leftHandSum := 0
	minDifference := -1
	//Get the total left hand sum
	for _, num := range someArray {
		rightHandSum += num
	}
	//start our loop from the last
	for index := 0; index < len(someArray); index++ {
		leftHandSum += someArray[index]
		rightHandSum -= someArray[index]
		tempDifference := int(math.Abs(float64(leftHandSum) - float64(rightHandSum)))

		if minDifference == -1 {
			minDifference = tempDifference
		}
		if minDifference == 0 {
			//Cant get lower than zero
			minDifference = tempDifference
			break
		}
		if minDifference > tempDifference {
			minDifference = tempDifference
		}
	}
	return minDifference
}
