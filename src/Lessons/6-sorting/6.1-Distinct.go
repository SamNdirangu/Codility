package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(solution([]int{}))
}

func solution(someArray []int) int {
	//sort the array
	sort.Ints(someArray)
	//loop and find distinct numbers
	disticntNumbers := 0
	for index := 0; index < len(someArray); index++ {
		if index == 0 {
			disticntNumbers++
		} else if someArray[index] != someArray[index-1] {
			disticntNumbers++
		}
	}
	return disticntNumbers

}
