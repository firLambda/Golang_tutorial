package main

import "fmt"

func integers() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	ints := integers()
	fmt.Println(ints())
	fmt.Println(ints())
	fmt.Println(ints())
	fmt.Println(ints())
	fmt.Println(ints())

	otherInts := integers()
	fmt.Println(otherInts())
	fmt.Println(otherInts())
	fmt.Println(otherInts())

}
