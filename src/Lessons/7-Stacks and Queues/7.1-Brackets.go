package main

import "fmt"

func main() {
	fmt.Println(solution(")("))
}

func solution(someString string) int {
	//check if S is empty
	if len(someString) == 0 {
		return 1
	}

	var bracketsTracker []string
	for index := 0; index < len(someString); index++ {
		character := string(someString[index])
		switch character {
		case ")":
			if len(bracketsTracker) < 1 || bracketsTracker[len(bracketsTracker)-1] != "(" {
				return 0
			}
			bracketsTracker = append(bracketsTracker[:len(bracketsTracker)-1])
		case "}":
			if len(bracketsTracker) < 1 || bracketsTracker[len(bracketsTracker)-1] != "{" {
				return 0
			}
			bracketsTracker = append(bracketsTracker[:len(bracketsTracker)-1])
		case "]":
			if len(bracketsTracker) < 1 || bracketsTracker[len(bracketsTracker)-1] != "[" {
				return 0
			}
			bracketsTracker = append(bracketsTracker[:len(bracketsTracker)-1])
		default:
			bracketsTracker = append(bracketsTracker, character)
		}
	}

	if len(bracketsTracker) != 0 {
		return 0
	}
	return 1 //if correctly matched brackets
}
