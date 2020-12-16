package tree

import "fmt"

/*
	请实现一个函数按照之字形顺序打印二叉树，即第一行按照从左到右的顺序打印，第二层按照从右到左的顺序打印，第三行再按照从左到右的顺序打印，其他行以此类推。
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
	2. cur出队，打印cur,基数行将cur子节点按左右顺序入队,偶数行将cur子节点按右左顺序入队
	3. 循环1，2操作，直到队列为空
	时间复杂度:O(nlogn)
	空间复杂度:O(n)
*/
func LevelOrderIII(root *TreeNode) [][]int {
	if nil == root {
		return nil
	}
	var result = make([][]int, 0)
	cur := root
	queue := []*TreeNode{cur}
	flag := 1
	// 外层 O(logN)
	for len(queue) != 0 {
		length := len(queue)
		row := make([]int, 0)
		// 内层O(n)
		for length != 0 {
			queue, cur = dqueue(queue)
			length--
			if flag % 2 != 0 {
				row = append(row, cur.Val)
			} else {
				row = append([]int{cur.Val}, row[0:]...) // 数组的头部插入
			}

			if cur.Left != nil {
				queue = enqueue(queue, cur.Left)
			}
			if cur.Right != nil {
				queue = enqueue(queue, cur.Right)
			}

		}
		fmt.Println("row:", row, "flag:", flag,"mod:", flag % 2)

		flag++ // 因为是在入队的时候进行判断，所以自加要放在上面，因为这里实际上要入的是第二行数据了

		result = append(result, row)
	}
	return result
}

/*
	采用双端队列
	基数层:头部出队，尾部入队
	偶数层:头部入队，尾部出队
*/
func LevelOrderIII_2(root *TreeNode) [][]int {
	if nil == root {
		return nil
	}
	var result = make([][]int, 0)
	cur := root
	queue := []*TreeNode{cur}
	flag := 1
	// 外层 O(logN)
	for len(queue) != 0 {
		length := len(queue)
		row := make([]int, 0)
		if flag %2 != 0 {
			// 基数层
			for length != 0 {
				queue, cur = dqueue(queue) // 头部出队
				length--
				row = append(row, cur.Val)
				if cur.Left != nil {
					queue = enqueue(queue, cur.Left) // 尾部入队
				}
				if cur.Right != nil {
					queue = enqueue(queue, cur.Right)
				}
			}
		} else {
			// 偶数层
			for length != 0 {
				queue, cur = queue[:len(queue)-1], queue[len(queue)-1]  // 尾部出队
				length--
				row = append(row, cur.Val)
				if cur.Right != nil {
					queue = enqueueFromHead(queue, cur.Right) // 头部入队
				}
				if cur.Left != nil {
					queue = enqueueFromHead(queue, cur.Left)
				}
			}
		}

		flag++ // 因为是在入队的时候进行判断，所以自加要放在上面，因为这里实际上要入的是第二行数据了
		result = append(result, row)
	}
	return result
}

/*
	打印奇数层： 从左向右 打印，先左后右 加入下层节点；
	若 deque 为空，说明向下无偶数层，则跳出；
	打印偶数层： 从右向左 打印，先右后左 加入下层节点；
*/