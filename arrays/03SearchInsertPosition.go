package main

import (
	"fmt"
)

func main() {
	f := fmt.Println
	array := []int{1, 3, 5, 6}
	f(searchInsert(array, 7))
}

func searchInsert(nums []int, target int) int {
	for i := 0; i < len(nums); i++ {
		if nums[i] == target {
			return i
		}
		if i+1 < len(nums) && nums[i] < target && nums[i+1] > target {
			return i + 1
		}
		if i+1 == len(nums) && nums[len(nums)-1] < target {
			return len(nums)
		}
	}
	return 0
}
