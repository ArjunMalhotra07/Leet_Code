package main

import (
	"github.com/ArjunMalhotra07/Leet_Code_GoLang/Binary_Search"
	"github.com/ArjunMalhotra07/Leet_Code_GoLang/Cyclic_Sort"
	"github.com/ArjunMalhotra07/Leet_Code_GoLang/Strings"
	"github.com/ArjunMalhotra07/Leet_Code_GoLang/arrays"
)

func main() {
	// mainDirectory()
	// arraysSubDirectory()
	binarySearchSubDirectory()
}

func mainDirectory() {
	// Main directory
	isPalindrome()
	findNumbersHelper()
	llHelperfunc()
	maximumWealthHelper()
	twoSumHelper()

}

func arraysSubDirectory() {
	// Arrays Sub Directory
	arrays.SingleNumberHelper()
	arrays.FloodFillHelper()
	arrays.IsToeplitzMatrixHelper()
	arrays.SearchInsertHelper()
	arrays.GetSumHelperFunc()
}

func binarySearchSubDirectory() {
	// Binary Search Sub Directory
	Binary_Search.AscHelper()
	Binary_Search.SearchRangeHelper()
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

func stringsSubDirectory() {
	Strings.CapitalizeTitle()
	Strings.CountUsingFieldsHelper()
	Strings.FirstPalindromeHelper()
	Strings.Check2Helper()
	Strings.IsAnagramHelper()
	Strings.NumJewelsInStonesHelper()
	Strings.CheckAnsHelper()
	Strings.LengthOfLastWordHelper()
	Strings.ReverseString([]byte{'e', 'f', 'g'})
	Strings.StrstrHelper()
	Strings.IsIsomorphicHelper()
	Strings.MapWords()
	Strings.Reverse()
}
