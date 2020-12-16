package tree

import "fmt"

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
返回 false 。


[1,2,2,3,null,null,3,4,null,null,4]
限制：

1 <= 树的结点个数 <= 10000
*/
/*
	设叶子节点高度为0，采用递归的方式计算以每个节点为根的左右子节点的高度，然后判断高度差
*/
func IsBalanced(root *TreeNode) bool {
	if nil == root {
		return true
	}
	left := height(root.Left)
	right := height(root.Right)
	fmt.Println("left:", left, "right:", right,"balance:", left - right)
	return left - right <= 1 && left - right >= -1 && IsBalanced(root.Left) && IsBalanced(root.Right) // todo 这里需要注意，不能只判断root节点，还需要递归判断每一个左子节点和右子节点
}

func height(node *TreeNode) int {
	// return
	if nil == node {
		return 0
	}
	// current logic

	// down
	return 1 + max(height(node.Left), height(node.Right))
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}



/*
	对二叉树进行后序遍历，在每次打印到根节点的时候就已经可以得到左右子树的高度了
	1. 判断左右子树是否平衡
	2. 获取左右子树的高度差，判断其绝对值是否不超过1
*/
func IsBalanced2(root *TreeNode) bool {
	if nil == root {
		return true
	}
	return  treeHeight2(root) != -1
}

func treeHeight(node *TreeNode) int {
	if nil == node {
		return 0
	}
	leftHeight := treeHeight(node.Left) // the height of left child tree
	rightHeight := treeHeight(node.Right) // the height of right child tree
	if leftHeight >= 0 && rightHeight >= 0 && leftHeight - rightHeight >= -1 && leftHeight - rightHeight <= 1 {
		return max(leftHeight, rightHeight) + 1
	} else { // not balance
		return -1
	}
}
// 剪枝，如果某一个节点为根的子树不平衡，直接返回，不再继续递归
func treeHeight2(node *TreeNode) int {
	if nil == node {
		return 0
	}
	leftHeight := treeHeight2(node.Left) // the height of left child tree

	if leftHeight == -1 {
		return -1
	}
	rightHeight := treeHeight2(node.Right) // the height of right child tree
	if rightHeight == -1 {
		return -1
	}
	if leftHeight - rightHeight > 1 || leftHeight - rightHeight < -1 {
		return -1
	}
	return max(leftHeight, rightHeight) + 1
}
