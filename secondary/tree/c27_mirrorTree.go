package tree

/*
	请完成一个函数，输入一个二叉树，该函数输出它的镜像。

	例如输入：

	     4
	   /   \
	  2     7
	 / \   / \
	1   3 6   9
	镜像输出：

	     4
	   /   \
	  7     2
	 / \   / \
	9   6 3   1
*/

// 要输出二叉树的镜像，其实就是左右子树对调，并一直递归
func mirrorTree(root *TreeNode) *TreeNode {
	if nil == root {
		return nil
	}
	root.Left, root.Right = root.Right, root.Left
	mirrorTree(root.Left)
	mirrorTree(root.Right)
	return root
}

/*
	BFS
*/
func mirrorTreeBFS(root *TreeNode) *TreeNode {
	if nil == root {
		return nil
	}
	queue := []*TreeNode{root}
	curNode := root
	for len(queue) != 0 {
		queue, curNode = dqueue(queue)
		curNode.Left, curNode.Right = curNode.Right, curNode.Left
		if curNode.Left != nil {
			queue = append(queue, curNode.Left)
		}
		if curNode.Right != nil {
			queue = append(queue, curNode.Right)
		}
	}
	return root
}