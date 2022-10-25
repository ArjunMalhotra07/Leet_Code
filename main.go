package main

import (
	"fmt"

	"github.com/ArjunMalhotra07/Leet_Code_GoLang/Binary_Search"
	"github.com/ArjunMalhotra07/Leet_Code_GoLang/Cyclic_Sort"
	"github.com/ArjunMalhotra07/Leet_Code_GoLang/arrays"
)

func main() {
	mainDirectory()
}
func mainDirectory() {
	f := fmt.Println
	isPalindrome()
	findNumbersHelper()
	llHelperfunc()
	matrix := [][]int{{1, 2, 3}, {3, 2, 1}}
	f(maximumWealth(matrix))
	f()
	f()
	twoSumHelper()

}

func arraysSubDirectory() {
	// Arrays Sub Directory
	arrays.SingleNumberHelper()
	arrays.FloodFillHelper()
	arrays.IsToeplitzMatrixHelper()
	arrays.SearchInsertHelper()
	arrays.MaxSubArray([]int{5, 10, 15, 20, 25})
	arrays.GetSumHelperFunc()
}

func binarySearchSubDirectory() {
	// Arrays Sub Directory
	Binary_Search.PeakIndexInMountainArrayHelper()

}

func cyclicSortSubDirectory() {
	Cyclic_Sort.SelectionSortHelperFunc()
	Cyclic_Sort.FindDisappearedNumbersHelper()
	Cyclic_Sort.SelectionSortFuncHelper()
	Cyclic_Sort.FindDuplicatesHelper()
	Cyclic_Sort.FindErrorNumsHelper()
	Cyclic_Sort.FirstMissingPositiveHelper()
}
