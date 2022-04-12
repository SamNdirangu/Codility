package main

import "fmt"

func main() {
	fmt.Println(solution([]int{1, 2, 3, 4, 5}, 0))
}

func solution(someArray []int, rotationTimes int) []int {
	//Rotate the array the given times
	//if rotation is zero or the same as the len of array no change will appear
	if rotationTimes == 0 || len(someArray) == rotationTimes {
		return someArray
	}
	//To avoid unnecessary long loop
	if rotationTimes > len(someArray) {
		//we'll rotate the remainder number of times
		rotationTimes = rotationTimes % len(someArray)
	}
	for index := 0; index < rotationTimes; index++ {
		//remove the last item and append it as the first
		rotateInt := []int{someArray[len(someArray)-1]} //convert to slice
		someArray = append(rotateInt, someArray[:len(someArray)-1]...)
	}
	return someArray
}
