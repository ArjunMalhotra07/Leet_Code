package trees

func MaxDepthUsingQ(root *TreeNode) int {
	q := &queue{}
	if root == nil {
		return 0
	}
	ans := 0
	q.enqueue(root)
	for len(q.nodes) != 0 {
		levelSize := len(q.nodes)
		for i := 0; i < levelSize; i++ {
			d := q.dequeue()
			if d.Left != nil {
				q.nodes = append(q.nodes, d.Left)
			}
			if d.Right != nil {
				q.nodes = append(q.nodes, d.Right)
			}
		}
		ans += 1
	}
	return ans
}

type queue struct {
	nodes []*TreeNode
}

func (q *queue) enqueue(node *TreeNode) {
	q.nodes = append(q.nodes, node)
}
func (q *queue) dequeue() *TreeNode {
	length := len(q.nodes)
	if length != 0 {
		removedVertex := q.nodes[0]
		if length != 1 {
			q.nodes = q.nodes[1:]
		} else {
			q.nodes = []*TreeNode{}
		}
		return removedVertex
	}
	return nil
}
