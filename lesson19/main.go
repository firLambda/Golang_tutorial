package main

import "fmt"

func main() {
	a := 0
	if a == 2 {
		fmt.Println("two")
	} else if a == 1{
		fmt.Println("one")
	} else{
		fmt.Println("I dont konw")
	}

	if b := 100; b == 100{
		fmt.Print("one hundred")
	}

	x := 2
	if x := 0; true{
		fmt.Println(x)
	}
	fmt.Println(x)
}