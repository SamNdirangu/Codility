package main

import "fmt"

func main() {

	someArray := []int{1, 2, 3, 4, 5, 6}

	someArray = append(someArray[:0])
	fmt.Println(someArray)

}
