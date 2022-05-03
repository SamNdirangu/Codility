package main

import (
	"fmt"
	"sync"
)

var waitGroup = sync.WaitGroup{}

func main() {
	fmt.Println(solution("GGGGGGGGGGGGGGGGGGGGGGGGTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGG", []int{2, 5, 0}, []int{30, 20, 10}))
}

func solution(S string, P, Q []int) []int {
	var impactFactor []int

	start, end := 0, 0
	//Loop 2: Loop through P and Q
	for index := 0; index < len(P); index++ {
		start = P[index]
		end = Q[index]
		waitGroup.Add(1)
		go func(start, end int) {
			factor := 5
		search:
			for i := start; i <= end; i++ {
				switch string(S[i]) {
				case "A":
					factor = 1
					break search
				case "C":
					factor = 2
					break
				case "G":
					if factor > 3 {
						factor = 3
					}
					break
				case "T":
					if factor > 4 {
						factor = 4
					}
					break
				default:
					break search
				}
			}
			impactFactor = append(impactFactor, factor)
			waitGroup.Done()
		}(start, end)
	}
	waitGroup.Wait()
	return impactFactor
}
