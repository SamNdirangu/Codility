package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(solution([]int{-3, 1, 2, -2, 5, 6}))
}

func solution(someArray []int) int {
	//Get the maximum product of any triplet
	//sort array
	sort.Ints(someArray)
	lastIndex := len(someArray) - 1
	//multiply the last 3 digits which are the largest
	product := someArray[lastIndex] * someArray[lastIndex-1] * someArray[lastIndex-2]
	return product
}
