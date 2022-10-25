package arrays

import "fmt"

func IsToeplitzMatrixHelper() {
	arr := [][]int{{1, 2, 3, 4}, {5, 1, 2, 3}, {9, 5, 1, 2}}
	arr2 := [][]int{{1, 2}, {2, 2}}
	loopExamples(arr)
	loopExamples(arr2)
}

func isToeplitzMatrix(matrix [][]int) bool {
	return helperFunc(matrix)
}
func helperFunc(testMatrix [][]int) bool {
	rows := len(testMatrix)
	columns := len(testMatrix[0])
	flag := true
	for i := 0; i < (rows); i++ {
		for j := 0; j < columns; j++ {
			if i+1 < rows && j+1 < columns {
				if testMatrix[i][j] != testMatrix[i+1][j+1] {
					flag = false
					break
				}
			}
		}
	}
	return flag
}

func printLoop(arr [][]int) {
	for i := 0; i < len(arr); i++ {
		fmt.Print(arr[i])
		fmt.Println()
	}
}

func loopExamples(arr [][]int) {
	f := fmt.Println
	f("Is 2D Array Toeplitz?")
	f()
	printLoop(arr)
	f()
	ansBoolean := isToeplitzMatrix(arr)
	f("*** ", ansBoolean, " ***")
	f()

}
