package main

import (
	"fmt"
	"sort"
)

func main() {

	fmt.Println(solution([]int{-2, -5, 0}))
}

func solution(someArray []int) int {
	//Return the smallest positive integer that does not occur in A
	missingInteger := 1
	//Sort our array
	sort.Ints(someArray)
	for index := 0; index < len(someArray); index++ {
		num := someArray[index]
		if num > 0 {
			if num != missingInteger {
				//Check if previous number matches the missing
				if index < 1 {
					break
				}
				prevNumber := someArray[index-1]
				if prevNumber == missingInteger {
					missingInteger++
				} else {
					//if not we've found the missing number break
					break
				}
			}
		}
	}
	//incase no missing integer was found
	if missingInteger == someArray[len(someArray)-1] {
		missingInteger++
	}

	return missingInteger
}
