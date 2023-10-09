package main

import "fmt"

func main() {
	nums := []int{1, 0, 1, 1, 0, 1}
	lastZero := -1
	ans := 0

	for k, v := range nums {
		if v == 0 {
			lastZero = k
		} else {
			tmp := k - lastZero
			if tmp > ans {
				ans = tmp
			}
		}
	}
	fmt.Println(ans)
}
