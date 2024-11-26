package main

import (
	"fmt"
	"strconv"
)

func main() {
	var i int = 1
	fl64 := float64(i)

	fmt.Printf("i = %T\n", i)
	fmt.Printf("fl64 = %T\n", fl64)

	i2 := int(fl64)
	fmt.Printf("i2 = %T\n", i2)

	var s string = "100"
	fmt.Printf("s = %T\n", s)

	i5, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(i5)
	fmt.Printf("i5 = %T\n", i5)

	var i6 int = 200
	s6 := strconv.Itoa(i6)
	fmt.Println(s6)
	fmt.Printf("s6 = %T\n", s6)

	var h string = "Hello World"
	b := []byte(h)
	fmt.Println(b)

	h2 := string(b)
	fmt.Print(h2)
}
