package main

import "fmt"

func fn(a []int) {
	a[2] = 5         // 0 1 5
	a = append(a, 6) // 0 1 5 6
	fmt.Println(a)   // 0 1 5 6
	a = append(a, 7) // 0 1 5 6 7
	a[0] = 5         // 5 1 5 6 7
	fmt.Println(a)   // 5 1 5 6 7
}

func main() {
	a := make([]int, 0, 5) // 0 1 2 3 4
	for i := 0; i < 4; i++ {
		a = append(a, i)
	}
	fn(a[:3])
	fmt.Println(a)
}
