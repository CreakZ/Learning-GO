package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	var num, positive int
	for i := 0; i < n; i++ {
		fmt.Scan(&num)
		if num > 0 {
			positive++
		}
	}
	fmt.Println(positive)
}
