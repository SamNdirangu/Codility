package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(solution(10, 40, 5))
}

func solution(X, Y, D int) int {
	//Count the minimal number of jumps that the small frog must perform to reach its target.
	//X is start Y is finish D is size of jump
	distance := Y - X
	numOfJumps := math.Ceil(float64(distance) / float64(D))

	return int(numOfJumps) //convert to int
}
