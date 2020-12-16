package tree

/*
	输入一棵二叉树的根节点，判断该树是不是平衡二叉树。如果某二叉树中任意节点的左右子树的深度相差不超过1，那么它就是一棵平衡二叉树。

	示例 1:

	给定二叉树 [3,9,20,null,null,15,7]

		3
	   / \
	  9  20
		/  \
	   15   7
	返回 true 。

	示例 2:

	给定二叉树 [1,2,2,3,3,null,null,4,4]

		   1
		  / \
		 2   2
		/ \
	   3   3
	  / \
	 4   4
	返回 false 。

*/

func isBalanced(root *TreeNode) bool {
	if nil == root {
		return true
	}
	height := getDepth(root.Left) - getDepth(root.Right)
	if height >= -1 && height <= 1{
		return isBalanced(root.Left)&&isBalanced(root.Right)
	} else {
		return false
	}
}

func getDepth(node *TreeNode) int {
	if nil == node {
		return 0
	}
	left := getDepth(node.Left)
	if left == -1 {  // 剪枝
		return -1
	}

	right := getDepth(node.Right)
	if right == -1 { // 剪枝
		return -1
	}
	if  left - right >= -1 && left - right <= 1 {
		return _max(getDepth(node.Left), getDepth(node.Right)) + 1
	}
	return -1
}

func isBalanced2(root *TreeNode) bool {
	return getDepth(root) != -1
}