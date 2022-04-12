package main

import (
	"fmt"
	"sort"
)

func main() {

	fmt.Println(solution([]int{1, 2, 11, 14, 3, 4, 5, 6, 7, 8, 9, 10, 12}))
}

func solution(someArray []int) int {
	if len(someArray) == 0 {
		return 1
	}
	missingNumber := 0
	//sort our array
	sort.Ints(someArray)
	for index := 0; index < len(someArray); index++ {
		number := index + 1
		if someArray[index] != number {
			missingNumber = number
			break
		}
	}
	//if no missing number, the missing number is the next largest number
	if missingNumber == 0 {
		missingNumber = someArray[len(someArray)-1] + 1
	}
	return missingNumber
}
