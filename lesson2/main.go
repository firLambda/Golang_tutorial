package main

import "fmt"

func HelloWorld() {
	s3 := "HelloWorld"
	fmt.Println(s3)
}

func main() {
	var i int = 100
	fmt.Println(i)

	var s string = "Hello World"
	fmt.Println(s)

	var t, f bool = true, false
	fmt.Println(t, f)

	var (
		i2 int    = 200
		s2 string = "Golang"
	)
	fmt.Println(i2, s2)

	var i3 int
	var s3 string

	fmt.Println(i3, s3)

	i3 = 300
	s3 = "go"
	fmt.Println(i3, s3)

	i4 := 400
	fmt.Println(i4)

	s4 := "tmpGo"
	fmt.Println(s4)

	HelloWorld()
}
