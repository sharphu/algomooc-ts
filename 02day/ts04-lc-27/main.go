package main

import "fmt"

func main() {
	nums := []int{0, 1, 2, 2, 3, 0, 4, 2}
	val := 2
	ind := 0

	for _, v := range nums {
		if v != val {
			nums[ind] = v
			ind++
		}
	}

	fmt.Println(nums)
	fmt.Println(ind)
}
