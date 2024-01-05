package graphs

func FindCircleNum(isConnected [][]int) int {
	ans := 0
	length := len(isConnected)
	if length == 0 {
		return ans
	}
	visitedArray := make([]int, length)
	myQueue := &Queue{}
	for i := 0; i < len(visitedArray); i++ {
		if visitedArray[i] == 0 {
			ans += 1
			visitedArray[i] = 1
			myQueue.Enqueue(i)
			for len(myQueue.nums) != 0 {
				dequeued := myQueue.Dequeue()
				for j := 0; j < len(isConnected[dequeued]); j++ {
					if isConnected[dequeued][j] == 1 && visitedArray[j] == 0 {
						visitedArray[j] = 1
						myQueue.Enqueue(j)
					}
				}
			}
		}
	}
	return ans
}

type Queue struct {
	nums []int
}

func (q *Queue) Enqueue(value int) {
	q.nums = append(q.nums, value)
}

func (q *Queue) Dequeue() int {
	if len(q.nums) > 0 {
		removedElement := q.nums[0]
		if len(q.nums) == 1 {
			q.nums = []int{}
		} else {
			q.nums = q.nums[1:]
		}
		return removedElement
	}
	return -1
}
