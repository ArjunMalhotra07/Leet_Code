package graphs

func FloodFill(image [][]int, sr int, sc int, color int) [][]int {
	oColor := image[sr][sc]
	if oColor == color {
		return image
	}
	FloodFillHelper(&image, sr, sc, color, oColor)
	return image
}
func FloodFillHelper(matrix *[][]int, row, col, nColor, oColor int) {
	if row < 0 || col < 0 || row >= len(*matrix) || col >= len((*matrix)[row]) || (*matrix)[row][col] != oColor {
		return
	}
	(*matrix)[row][col] = nColor
	//! Left
	FloodFillHelper(matrix, row, col-1, nColor, oColor)
	//! Right
	FloodFillHelper(matrix, row, col+1, nColor, oColor)
	//! Top
	FloodFillHelper(matrix, row-1, col, nColor, oColor)
	//! Bottom
	FloodFillHelper(matrix, row+1, col, nColor, oColor)

}
