package main

import "fmt"

func createdRange() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	// i := 0
	// for {
	// 	i++
	// 	if i == 3 {
	// 		break
	// 	}
	// 	fmt.Println("Loop")
	// }

	// f_range := createdRange()
	// for i := f_range(); i < 10; i = f_range() {
	// 	fmt.Println(i)
	// }

	// for i := 0; i < 10; i++{
	// 	if i == 3{
	// 		continue
	// 	}
	// 	fmt.Println(i)
	// }

	// arr := [3]int{1,2,3}
	// for i := 0; i < len(arr); i++{
	// 	fmt.Println(arr[i])
	// }

	arr := [3]int{1, 2, 3}
	for i := range arr {
		fmt.Println(i)
	}

	sl := []string{"Python", "PHP", "GO"}
	for i, v := range sl {
		fmt.Println(i, v)
	}

	m := map[string]int{"apple": 100, "bananna": 300}
	for key, v := range m {
		fmt.Printf("%s:%d\n", key, v)
	}

}
