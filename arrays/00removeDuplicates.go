package main

import (
	"fmt"
)

func main() {
	array1 := []int{4, 1, 2, 1, 2, 10, 10, 10, 20}
	f := fmt.Println
	f(singleNumber(array1))

}

func singleNumber(nums []int) []int {
	return helper(nums)

}

func helper(nums []int) []int {
	ansArray := []int{}

	// declare Hashmap of input array
	countMap := make(map[int]int)
	length := len(nums)
	count := 1

	// Making of hashmap
	for i := 0; i < length-1; i++ {
		// If element is found already in the map, ignore move forward
		if _, found := countMap[nums[i]]; !found {
			for j := i + 1; j < length; j++ {
				if nums[i] == nums[j] {
					count++
				}
			}
			//giving value to key
			countMap[nums[i]] = count
			count = 1
		}
	}
	// Check if last element is in the Hashmap or not, if not then add it
	if _, found := countMap[nums[length-1]]; !found {
		countMap[nums[length-1]] = 1
	}
	fmt.Println(countMap)

	for ans, count := range countMap {
		if count == 1 {
			ansArray = append(ansArray, ans)
		}
	}
	return ansArray
}
