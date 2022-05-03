package main

import "fmt"

func main() {
	fmt.Println(Solution(")"))
}

func Solution(someString string) int {
	if len(someString) == 0 {
		//string is properly nested
		return 1
	}

	var bracketsTracker []string

	for index := 0; index < len(someString); index++ {
		character := string(someString[index])
		lastIndex := len(bracketsTracker) - 1
		switch character {
		case ")":
			if len(bracketsTracker) == 0 || bracketsTracker[lastIndex] != "(" {
				//incorrectly nested
				return 0
			}
			bracketsTracker = append(bracketsTracker[:lastIndex])
		default:
			bracketsTracker = append(bracketsTracker, character)
		}
	}
	if len(bracketsTracker) != 0 {
		return 0
	}
	//String is properly nested
	return 1
}
