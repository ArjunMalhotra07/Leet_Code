package main

import "fmt"

func main() {

	arr := [][]int{{0, 0, 0}, {0, 1, 1}}
	arr2 := [][]int{{1, 1, 1, 1}, {0, 1, 0, 0}, {0, 1, 0, 1}}
	loopExamples(arr)
	loopExamples(arr2)

}

func changeColor(arr [][]int, sr, sc, newColor int, column, row, presentColor int) [][]int {
	// fmt.Println(presentColor)
	if arr[sr][sc] == presentColor {
		arr[sr][sc] = newColor
	}
	if sc-1 != -1 {
		if arr[sr][sc-1] == presentColor {
			changeColor(arr, sr, sc-1, newColor, column, row, presentColor)
		}
	}
	if sr-1 != -1 {
		if arr[sr-1][sc] == presentColor {
			changeColor(arr, sr-1, sc, newColor, column, row, presentColor)
		}
	}
	if sr+1 < row {
		if arr[sr+1][sc] == presentColor {
			changeColor(arr, sr+1, sc, newColor, column, row, presentColor)
		}
	}
	if sc+1 < column {
		if arr[sr][sc+1] == presentColor {
			changeColor(arr, sr, sc+1, newColor, column, row, presentColor)
		}
	}

	return (arr)
}
func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	presentColor := image[sr][sc]
	if presentColor == newColor {
		return image
	}
	column := len(image[0])
	row := len(image)
	return changeColor(image, sr, sc, newColor, column, row, presentColor)
}

func printLoop(arr [][]int) {
	for i := 0; i < len(arr); i++ {
		fmt.Print(arr[i])
		fmt.Println()
	}
}

func loopExamples(arr [][]int) {
	f := fmt.Println
	f("*********")
	f("Input Array")
	printLoop(arr)
	f()
	ansArray := floodFill(arr, 1, 1, 2)
	f("Altered Array")
	printLoop(ansArray)
	f()

}
