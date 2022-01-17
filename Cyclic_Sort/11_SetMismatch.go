package main

import "fmt"

func main() {
	arr := []int{1, 1}
	ans := findErrorNums(arr)
	fmt.Println(ans)

}

func findErrorNums(nums []int) []int {
	arr := []int{}

	i := 0
	for i < len(nums) {
		correct := nums[i] - 1
		if nums[i] != nums[correct] {
			swap(nums, i, correct)
		} else {
			i++
		}
	}

	for i := 0; i < len(nums); i++ {
		if nums[i] != i+1 {
			arr = append(arr, nums[i], i+1)
		}
	}

	return arr
}
func swap(nums []int, i int, correct int) {
	temp := nums[i]
	nums[i] = nums[correct]
	nums[correct] = temp
}
