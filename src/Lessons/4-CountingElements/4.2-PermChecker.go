package main

import "fmt"

func main() {

	fmt.Println(solution([]int{4, 5, 2, 3}))
	//expect 0
}

func solution(someArray []int) int {
	//Check if array is a permutation
	//Calculate the total sum of ints if permutation
	maxInt := len(someArray)
	expectedSum := ((1 + maxInt) * maxInt) / 2 //Formulae for series sum of  1 + n + ...n+1
	permTracker := make([]bool, maxInt)

	currentSum := 0
	for index := 0; index < len(someArray); index++ {
		num := someArray[index]
		if num > maxInt {
			// a correct permutation will not have a number larger than the largest expected number
			break
		}
		if !permTracker[num-1] {
			permTracker[num-1] = true
			currentSum += num
		}
		if currentSum > expectedSum {
			break
		}
	}
	if currentSum != expectedSum {
		return 0
	}
	return 1

}
