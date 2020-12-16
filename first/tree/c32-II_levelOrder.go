package tree

import "fmt"

/*
	从上到下按层打印二叉树，同一层的节点按从左到右的顺序打印，每一层打印到一行。
	例如:
	给定二叉树: [3,9,20,null,null,15,7],

		3
	   / \
	  9  20
		/  \
	   15   7
	返回其层次遍历结果：

	[
	  [3],
	  [9,20],
	  [15,7]
	]

	提示：
	节点总数 <= 1000

*/

/*
	队列
	设定当前节点为cur
	1. cur入队
	2. cur出队，打印cur,将cur子节点按左右顺序入队
	3. 循环1，2操作，直到队列为空
	时间复杂度:O(n)
	空间复杂度:O(n)
*/
func LevelOrder(root *TreeNode) [][]int {
	if nil == root {
		return nil
	}
	var result  = make([][]int, 0)
	cur := root
	queue := []*TreeNode{cur}
	for len(queue) != 0 {
		length := len(queue)
		row := make([]int, 0)
		for length != 0 {
			queue, cur = dqueue(queue)
			length--
			//if nil == cur { // 这里不需要，在拿到当前queue的最后一个节点之后，length会变为0，不会再进入内部循环
			//	continue
			//}
			row = append(row, cur.Val)
			if cur.Left != nil {
				queue = enqueue(queue,cur.Left)
			}
			if cur.Right != nil {
				queue = enqueue(queue,cur.Right)
			}
		}
		result = append(result, row)
	}
	return result
}

func LevelOrder1(root *TreeNode) [][]int {
	if nil == root {
		return nil
	}
	var result  = make([][]int, 0)
	cur := root
	queue := []*TreeNode{cur}
	for len(queue) != 0 {
		length := len(queue)
		row := make([]int, 0)
		for length != 0 {
			queue, cur = dqueue(queue)
			length--
			if nil == cur { // 这里不需要，在拿到当前queue的最后一个节点之后，length会变为0，不会再进入内部循环
				continue
			}
			row = append(row, cur.Val)

			queue = enqueue(queue,cur.Left)
			queue = enqueue(queue,cur.Right)
		}
		fmt.Println("row", row,"isNil",row == nil, "length", len(row))
		if len(row) != 0 {
			result = append(result, row)
		}

	}
	return result
}
