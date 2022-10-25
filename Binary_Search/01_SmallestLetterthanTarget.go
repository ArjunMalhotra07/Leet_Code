//Ascending order: https://leetcode.com/problems/find-smallest-letter-greater-than-target/submissions/

package Binary_Search

import "fmt"

func AscHelper() {
	f := fmt.Println
	f("PROBLEM 1 : To find smallest and largest rune compared to 'a' in []byte{'c', 'f', 'j'}")
	f("Smallest than 'a' in array -- ", Asc([]byte{'c', 'f', 'j'}, 'a'))
	f("Largest than 'a' in array -- ", Desc([]byte{'j', 'f', 'c'}, 'a'))
	f()
	f()
}
func Asc(arr []byte, target byte) string {
	start := 0
	end := len(arr) - 1
	ansArray := []byte{}

	for start <= end {
		mid := start + (end-start)/2
		if target < arr[mid] {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}
	ansArray = append(ansArray, arr[start%len(arr)])
	return string(ansArray)
}

// Decreasing order vaste
func Desc(arr []byte, target byte) string {
	start := 0
	end := len(arr) - 1
	ansArray := []byte{}

	for start <= end {
		mid := start + (end-start)/2
		if arr[mid] > target {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}
	ansArray = append(ansArray, arr[start%len(arr)])
	return string(ansArray)
	// return (arr[start%len(arr)])
}
