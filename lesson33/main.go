package main

import (
	. "fmt"
	"golang_udemy/lesson33/foo"
)

func appName() string {
	const AppName = "GoApp"
	var Version string = "1.0"
	return AppName + " " + Version
}

func Do(s string) (b string) {
	// var b string = s
	
	b = s
	{
		b := "BBBB"
		Println(b)
	}
	Println(b)
	return b
}

func main() {
	Println(foo.Max)
	Println(foo.ReturnMin())

	Println(appName())
	Println(Do("AAA"))
}
