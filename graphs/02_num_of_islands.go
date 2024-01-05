package graphs

// https://leetcode.com/problems/number-of-islands/description/
func NumIslands(grid [][]byte) int {
	count := 0
	for row := 0; row < len(grid); row++ {
		for col, cell := range grid[row] {
			if cell == '1' {
				count++
				BfS(&grid, row, col)
			}
		}
	}
	return count
}
func BfS(grid *[][]byte, i int, j int) {
	if i < 0 || j < 0 || i >= len(*grid) || j >= len((*grid)[i]) || (*grid)[i][j] == '0' {
		return
	}
	(*grid)[i][j] = '0'
	BfS(grid, i-1, j)
	BfS(grid, i+1, j)
	BfS(grid, i, j+1)
	BfS(grid, i, j-1)
}
