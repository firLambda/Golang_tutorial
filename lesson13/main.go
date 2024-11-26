package main

import "fmt"

func Plus(x int, y int) int {
	return x + y
}

func Div(x, y int) (int, int) {
	q := x / y
	r := x % y
	return q, r
}

func Double(price int)(result int){
	result = price * 2
	return 
}

func Noreturn(){
	fmt.Println("No Return")
	return
}

func main() {
	i := Plus(1, 2)
	fmt.Println(i)

	i2, i3 := Div(4,3)
	fmt.Println(i2, i3)

	result := Double(100)
	fmt.Println(result)

	Noreturn()
}
