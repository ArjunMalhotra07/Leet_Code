//https://leetcode.com/problems/find-first-and-last-position-of-element-in-sorted-array/submissions/

func searchRange(nums []int, target int) []int {
	start := search(nums, target, true)
	end := search(nums, target, false)

	return []int{start, end}

}

func search(nums []int, target int, findStartIndex bool) int {
	ans := -1
	start := 0
	end := len(nums) - 1

	for start <= end {

		mid := start + (end-start)/2
		if nums[mid] > target {
			end = mid - 1
		} else if nums[mid] < target {
			start = mid + 1
		} else {
			ans = mid
			if findStartIndex == true {
				end = mid - 1
			} else {
				start = mid + 1
			}
		}
	}
	return ans
}

