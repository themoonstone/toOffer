package tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

/*
	输入一棵二叉树的根节点，求该树的深度。从根节点到叶节点依次经过的节点（含根、叶节点）形成树的一条路径，最长路径的长度为树的深度。

	例如：

	给定二叉树 [3,9,20,null,null,15,7]，

		3
	   / \
	  9  20
		/  \
	   15   7
	返回它的最大深度 3 。
*/

/*
	recursion
	当前节点的最大深度等于其左子树的深度与右子树的深度的最大值+1
*/
func MaxDepth(root *TreeNode) int {
	if nil == root {
		return 0
	}
	return _max(MaxDepth(root.Left), MaxDepth(root.Right)) + 1
}


/*
	BFS:
*/