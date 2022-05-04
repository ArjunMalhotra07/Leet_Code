package main

import "fmt"

func main() {

	arr := [][]int{{0, 0, 0}, {0, 1, 1}}
	f := fmt.Println
	loop(arr)
	f()
	ansArray := floodFill(arr, 1, 1, 1)
	loop(ansArray)
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
	fmt.Println(presentColor, column, row)
	return changeColor(image, sr, sc, newColor, column, row, presentColor)
}

func loop(arr [][]int) {
	for i := 0; i < len(arr); i++ {
		fmt.Print(arr[i])
		fmt.Println()
	}
}
