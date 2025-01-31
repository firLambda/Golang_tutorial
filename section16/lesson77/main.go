package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.OpenFile("foo.txt", os.O_RDONLY, 0666)
	if err != nil {
		log.Fatalln(err)
	}

	defer f.Close()

	bs := make([]byte, 120)
	n, _ := f.Read(bs)
	fmt.Println(n)
	fmt.Println(string(bs))
}
