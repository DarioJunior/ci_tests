package main

import "fmt"

func main() {
	fmt.Println(Sum(5, 5))
}

// Sum is a function that takes two integers and returns the sum of them
func Sum(a int, b int) int {
	return a + b
}