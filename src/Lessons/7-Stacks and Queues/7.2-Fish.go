package main

import "fmt"

func main() {

	fmt.Println(Solution([]int{4, 3, 2, 1, 5, 7}, []int{0, 1, 0, 0, 1, 0}))
}

func Solution(A, B []int) int {
	// A is fish number all distinct
	// B is fish direction 0 upstream 1 downstream

	//if two fish moving against each other the bigger one survives
	//return the number of fish that will stay alive
	eaterFish := 0
	aliveFishes := len(A) //All fish are alive
	downstreamFishes := []int{}

	for index := 0; index < len(A); index++ {
		fishNum := A[index]
		isDownStream := B[index] == 1
		lastDownStreamFish := len(downstreamFishes) - 1

		if isDownStream {
			downstreamFishes = append(downstreamFishes, fishNum)
		} else {
			if fishNum < downstreamFishes[lastDownStreamFish] {
				//fish gets eaten
				aliveFishes--
			} else {
				if eaterFish != 0 {
					//eater fish got eaten
					aliveFishes--
				}
				eaterFish = 0
			}
		}
	}
	return aliveFishes
}
