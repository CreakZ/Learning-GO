package main

import (
	"fmt"
)

func main() {
	var a, b, n int
	fmt.Scan(&a, &b)

	t := a
	for t > 0 {
		n++
		t /= 10
	}

	for n > 0 {
		var ten = 10
		for j := 0; j < n-2; j++ {
			ten *= 10
		}
		var num1, temp int
		if a > 9 {
			num1 = a / ten
		} else {
			num1 = a
		}
		temp = b
		for temp > 0 {
			num2 := temp % 10
			if num1 == num2 {
				fmt.Print(num1, " ")
			}
			temp /= 10
		}
		a %= ten
		n--
	}
}
