package graphs

func updateMatrix(mat [][]int) [][]int {
	myQueue := &MatrixQueue{}
	visited := make([][]int, len(mat))
	distMatrix := make([][]int, len(mat))
	//! Initialising Visited and Distance Matrix
	for i := 0; i < len(mat); i++ {
		visited[i] = make([]int, len(mat[i]))
		distMatrix[i] = make([]int, len(mat[i]))
	}
	//! Enque-ing all the points with Value as 0
	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat[i]); j++ {
			if mat[i][j] == 0 {
				myQueue.Enqueue([]int{i, j, 0})
				visited[i][j] = 1
				distMatrix[i][j] = 0
			}
		}
	}
	//! Loop through the Queue and enqueue all the points which are not visited Updating distance matrix simult.
	for len(myQueue.nums) != 0 {
		dequeued := myQueue.Dequeue()
		row := dequeued[0]
		col := dequeued[1]
		steps := dequeued[2]
		distMatrix[row][col] = steps
		nRow := []int{1, -1, 0, 0}
		nCol := []int{0, 0, -1, 1}
		//! Going in all 4 directions: down, up, left, right
		for i := 0; i < 4; i++ {
			nrow := row + nRow[i]
			ncol := col + nCol[i]
			if nrow >= 0 && ncol >= 0 && nrow < len(mat) && ncol < len(mat[nrow]) && visited[nrow][ncol] != 1 {
				myQueue.Enqueue([]int{nrow, ncol, steps + 1})
				visited[nrow][ncol] = 1
			}
		}

	}
	return distMatrix
}

type MatrixQueue struct {
	nums [][]int
}

func (q *MatrixQueue) Enqueue(value []int) {
	q.nums = append(q.nums, value)
}

func (q *MatrixQueue) Dequeue() []int {
	if len(q.nums) > 0 {
		removedElement := q.nums[0]
		if len(q.nums) == 1 {
			q.nums = [][]int{}
		} else {
			q.nums = q.nums[1:]
		}
		return removedElement
	}
	return []int{}
}
