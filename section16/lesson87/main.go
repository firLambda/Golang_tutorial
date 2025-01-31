package main

import (
	"fmt"
	"strings"
)

func main() {
	s1 := strings.Join([]string{"A", "B", "C"}, ",")
	s2 := strings.Join([]string{"A", "B", "C"}, "")
	fmt.Println(s1, s2)

	s3 := strings.Repeat("ABC", 4)
	s4 := strings.Repeat("ABC", 0)

	fmt.Println(s3, s4)

	s5 := strings.Replace("AAAAAAA", "A", "B", 1)
	s6 := strings.Replace("AAAAAAA", "A", "B", -1)

	fmt.Println(s5, s6)

}
