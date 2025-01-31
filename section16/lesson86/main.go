package main

import (
	"fmt"
	"strconv"
)

func main() {
	bt := true
	fmt.Printf("%v\n", strconv.FormatBool(bt))

	i := strconv.FormatInt(-100, 2)
	fmt.Printf("%v, %T\n", i, i)
	i2 := strconv.Itoa(100)
	fmt.Printf("%v, %T\n", i2, i2)
}
