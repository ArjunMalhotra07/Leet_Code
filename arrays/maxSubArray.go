package main

import "fmt"

func main() {
	// array1 := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	array1 := []int{-2, -1, -5}
	getSum(array1)

}

func getSum(nums []int) {
	helperFunction(nums)

}

func helperFunction(nums []int) {
	var ansSum int

	f := fmt.Println
	if len(nums) == 1 {
		f(nums[0])
	}
	testSubArray := [][]int{}
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			if j+1 > i {
				testSubArray = append(testSubArray, nums[i:j+1])

			}
		}

	}
	for i := 0; i < len(testSubArray); i++ {
		var sum int

		f(testSubArray[i])
		for j := 0; j < len(testSubArray[i]); j++ {
			sum = sum + testSubArray[i][j]
		}
		// testSum = sum
		// ansSum = testSum
		if i == 0 {
			ansSum = sum
		}
		if sum > ansSum {
			ansSum = sum
		}
		f(sum)
		f()

		sum = 0

	}
	f(ansSum)
}
