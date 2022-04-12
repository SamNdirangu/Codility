package main

import (
	"fmt"
	"sort"
)

func main() {
	//Return 10
	fmt.Println(solution([]int{2, 2, 4, 4, 5, 6, 6, 5, 7, 8, 9, 10, 7, 8, 9}))

}

func solution(someArray []int) int {
	//get the unpaired element
	if len(someArray) == 1 {
		//if only one item in array
		return someArray[0]
	}
	//sort our array
	sort.Ints(someArray)
	//loop through
	oddNumber := 0
	//since numbers are in pairs we loop increasing our index by 2
	for index := 1; index < len(someArray); index += 2 {
		number := someArray[index]
		previousNumber := someArray[index-1]
		if number != previousNumber {
			oddNumber = previousNumber
			break
		}
	}
	// if no oddnumber the oddnumber is then the last number in the array
	if oddNumber == 0 {
		oddNumber = someArray[len(someArray)-1]
	}
	return oddNumber
}
