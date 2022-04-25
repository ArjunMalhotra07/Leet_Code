package main

import (
	"fmt"
)

func main() {
	array1 := []int{4, 1, 2, 1, 2}
	f := fmt.Println
	f(singleNumber(array1))

}

func singleNumber(nums []int) int {
	return helper(nums)

}

func helper(nums []int) int {
	countMap := make(map[int]int)
	length := len(nums)
	flag := false
	for i := 0; i < length-1; i++ {
		if _, found := countMap[nums[i]]; !found {
			for j := i + 1; j < length; j++ {
				if nums[i] == nums[j] {
					countMap[nums[i]] = 2
					flag = true
				}
			}
			if !flag {
				countMap[nums[i]] = 1
			}
			flag = false
		}
	}
	if _, found := countMap[nums[length-1]]; !found {
		countMap[nums[length-1]] = 1
	}
	fmt.Println(countMap)

	for ans, count := range countMap {
		if count == 1 {
			return ans
		}
	}
	return 0
}
