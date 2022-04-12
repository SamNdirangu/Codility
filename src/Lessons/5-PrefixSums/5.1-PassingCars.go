package main

import "fmt"

func main() {
	fmt.Println(solution([]int{0, 1, 0, 1, 1}))
}
func solution(someArray []int) int {
	passingCars := 0
	pairs := 0

	for index := 0; index < len(someArray); index++ {
		if someArray[index] == 0 {
			passingCars++
		} else {
			pairs += passingCars
		}
		if pairs > 1000000000 {
			pairs = -1
			break
		}
		//fmt.Println("Car:: ", someArray[index])
		//fmt.Println("pcars:: ", passingCars, " pairs:: ", pairs)
	}
	return pairs
}
