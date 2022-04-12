package main

import "fmt"

func main() {

	fmt.Println(solution(5, []int{3, 4, 4, 6, 1, 4, 4}))
	//expect  [3, 2, 2, 4, 2]
}

func solution(numOfCounters int, someArray []int) []int {
	counters := make([]int, numOfCounters)
	counterMax := 0
	counterMaxTracker := 0

	//Loop through array
	for index := 0; index < len(someArray); index++ {
		num := someArray[index]
		counterIndex := num - 1

		if num >= 1 && num <= numOfCounters {
			//increase counter num by 1
			if counters[counterIndex] < counterMax {
				counters[counterIndex] = counterMax + 1
			} else {
				counters[counterIndex] += 1
			}
			if counters[counterIndex] > counterMaxTracker {
				counterMaxTracker = counters[counterIndex]
			}
		}
		if num == numOfCounters+1 {
			//max all counters by storing in counterMax
			counterMax = counterMaxTracker
		}
	}
	//Loop through our counters
	for index := 0; index < len(counters); index++ {
		counter := counters[index]
		//Max all counters less than the required min of countermax
		if counter < counterMax {
			counters[index] = counterMax
		}
	}

	return counters

}
