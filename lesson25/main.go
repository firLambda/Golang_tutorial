package main

import (
	"fmt"
	"os"
)

func TestDefer() {
	defer fmt.Println("END")
	fmt.Println("START")
}

func RunDefer() {
	defer fmt.Println("1")
	defer fmt.Println("2")
	defer fmt.Println("3")

}

func main() {
	TestDefer()

	RunDefer()

	file, err := os.Create("test.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	fmt.Println([]byte("Hello"))
	file.Write([]byte("Hello"))
}
