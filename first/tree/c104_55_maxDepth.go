package tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
	广度优先算法
*/
func MaxDepth(root *TreeNode) int {
	if nil == root {
		return 0
	}
	queue := []*TreeNode{root} // 初始化一个队列
	var cur *TreeNode
	var count int = 0
	for len(queue) > 0 {
		length := len(queue)
		for length > 0 {
			queue, cur = dequeue(queue)
			if cur != nil && cur.Left != nil {
				queue = append(queue, cur.Left)
				//enqueue(queue, cur.Left)
			}
			if cur != nil && cur.Right != nil {
				queue = append(queue, cur.Right)
				//enqueue(queue, cur.Right)
			}
			length--
		}
		count++
	}
	return count
}

