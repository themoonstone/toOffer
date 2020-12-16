package tree

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

	[3,9,20,15,7]

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
func LevelOrderII(root *TreeNode) []int {
	if nil == root {
		return nil
	}
	cur := root
	queue := []*TreeNode{cur}
	row := make([]int, 0)
	for len(queue) != 0 {
		length := len(queue)
		for length != 0 {
			queue, cur = dqueue(queue)
			length--
			row = append(row, cur.Val)
			if cur.Left != nil {
				queue = enqueue(queue,cur.Left)
			}
			if cur.Right != nil {
				queue = enqueue(queue,cur.Right)
			}
		}
	}
	return row
}