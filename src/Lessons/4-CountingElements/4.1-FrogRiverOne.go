package main

import "fmt"

func main() {

	fmt.Println(solution(3, []int{1, 3, 1, 3, 1, 2, 1, 3}))
}

func solution(position int, someArray []int) int {
	seconds := -1
	//create array to track steps taken
	//Get the total sum of required steps from 1 to position
	stepsRequired := ((1 + position) * position) / 2 //is a series
	stepsTaken := make([]bool, position)             //create array to store if a step is taken

	for index := 0; index < len(someArray); index++ {
		currentPosition := someArray[index]
		if currentPosition <= position && !stepsTaken[currentPosition-1] {
			stepsRequired -= currentPosition
			stepsTaken[currentPosition-1] = true
		}
		if stepsRequired == 0 {
			seconds = index
			break
		}
	}
	return seconds
}
