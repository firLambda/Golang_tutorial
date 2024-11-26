package main

import "fmt"

func Plus(x int, y int) int {
	return x + y
}

func main() {
	i := Plus(1, 2)
	fmt.Println(i)
}
