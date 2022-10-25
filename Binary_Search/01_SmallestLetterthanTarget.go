//Ascending order: https://leetcode.com/problems/find-smallest-letter-greater-than-target/submissions/

package Binary_Search

func Asc(arr []byte, target byte) byte {
	start := 0
	end := len(arr) - 1

	for start <= end {
		mid := start + (end-start)/2
		if target < arr[mid] {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}
	return arr[start%len(arr)]
}

// Decreasing order vaste
func Desc(arr []byte, target byte) byte {
	start := 0
	end := len(arr) - 1

	for start <= end {
		mid := start + (end-start)/2
		if arr[mid] > target {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}
	return arr[start%len(arr)]
}
