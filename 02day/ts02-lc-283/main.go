package main

import "fmt"

func main() {
	nums := []int{0, 3, 2, 5, 0, 7}

	slow := 0
	for k, v := range nums {
		if v != 0 {
			nums[slow] = nums[k]
			slow++
		}
	}
	for i := slow; i < len(nums); i++ {
		nums[i] = 0
	}

	fmt.Println(nums)
}
