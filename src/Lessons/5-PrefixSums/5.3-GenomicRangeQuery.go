package main

import (
	"fmt"
)

func main() {
	fmt.Println(solution("GGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGG", []int{100, 100, 0}, []int{200, 200, 200}))
}

func solution(S string, P, Q []int) []int {
	var positionA []int
	var positionC []int
	var positionG []int
	var positionT []int
	var impactFactor []int

	//Loop 1: through our dna String
	for index, letter := range S {
		switch string(letter) {
		case "A":
			positionA = append(positionA, index)
		case "C":
			positionC = append(positionC, index)
		case "G":
			positionG = append(positionG, index)
		case "T":
			positionT = append(positionT, index)
		}
	}
	factor := 5
	loopCount, start, end := 0, 0, 0
	testA, testC, testG, testT := false, false, false, false
	//Loop 2: Loop through P and Q
	for index := 0; index < len(P); index++ {
		factor = 5
		start = P[index]
		end = Q[index]
	search:
		for i := 0; i >= 0; i++ {
			fmt.Println(loopCount)
			loopCount++
			testA = len(positionA) > i
			testC = len(positionC) > i
			testG = len(positionG) > i
			testT = len(positionT) > i
			switch {
			case testA:
				if testA {
					if positionA[i] >= start && positionA[i] <= end {
						factor = 1
						break search
					}
				}
				fallthrough
			case testC:
				if testC {
					if positionC[i] >= start && positionC[i] <= end {
						if factor > 2 {
							factor = 2
						}
						if !testA {
							break search
						}
						break
					}
				}
				fallthrough
			case testG:
				if testG {
					if positionG[i] >= start && positionG[i] <= end {
						if factor > 3 {
							factor = 3
						}
						if !testA && !testC {
							break search
						}
						break
					}
				}
				fallthrough
			case testT:
				if testT {
					if positionT[i] >= start && positionT[i] <= end {
						if factor == 5 {
							factor = 4
						}
						if !testA && !testC && !testG {
							break search
						}
					}
				}
				break
			default:
				break search
			}
		}
		impactFactor = append(impactFactor, factor)
	}
	fmt.Println(loopCount)
	return impactFactor
}
