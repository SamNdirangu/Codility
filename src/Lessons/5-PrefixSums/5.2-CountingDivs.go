package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(solution(2, 9, 3))
}
func solution(A, B, K int) int {
	//Return integers between A and B that are divisible by K
	divisibleNumbers := 0
	if A%K == 0 {
		divisibleNumbers++
		divisibleNumbers += int(math.Floor(float64((B - A) / K)))
	} else {
		num := A + 1
		keepSearching := true
		for keepSearching == true {
			if num%K == 0 {
				divisibleNumbers++
				divisibleNumbers += int(math.Floor(float64((B - num) / K)))
				keepSearching = false
				break
			}
			num++
		}
	}
	return divisibleNumbers
}
