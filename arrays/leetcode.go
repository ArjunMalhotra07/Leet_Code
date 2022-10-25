package arrays

func MaxSubArray(nums []int) int {
	return getSumFunction(nums)
}

func getSumFunction(nums []int) int {
	return getSumHelperFunction(nums)

}

func getSumHelperFunction(nums []int) int {
	var ansSum int
	sum := 0

	if len(nums) == 1 {
		return nums[0]
	}
	testArray := []int{}

	testSubArray := [][]int{}
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			if j+1 > i {
				testSubArray = append(testSubArray, nums[i:j+1])
				testArray = nums[i : j+1]

				tempLength := len(testArray)
				for j := 0; j < tempLength; j++ {

					sum = sum + testArray[j]

				}

				if j == 0 {
					ansSum = sum
				}
				if sum > ansSum {
					ansSum = sum
				}

				sum = 0

			}
		}

	}

	return ansSum
}
