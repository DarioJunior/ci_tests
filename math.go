package main

import "fmt"

func main() {
	fmt.Println(sum(5, 5))
}

// Sum is a function that takes two integers and returns the sum of them
func sum(a int, b int) int {
	return a + b
}