package tree

/*
	请实现一个函数，用来判断一棵二叉树是不是对称的。如果一棵二叉树和它的镜像一样，那么它是对称的。

	例如，二叉树 [1,2,2,3,4,4,3] 是对称的。

	    1
	   / \
	  2   2
	 / \ / \
	3  4 4  3
	但是下面这个 [1,2,2,null,3,null,3] 则不是镜像对称的:

	    1						1
	   / \					   / \
	  2   2					  2	  2
	   \   \				 /	 /
	   3    3				3	3

*/

/*
	递归:
	这里要注意一个特点，二叉树的镜像表示将二叉树进行翻转，如果是对称的，则对称节点的value一定相等
*/
func isSymmetric(root *TreeNode) bool {
	if nil == root {
		return true
	}
	return _isSymmetric(root.Left, root.Right)

}

func _isSymmetric(left *TreeNode, right *TreeNode) bool {
	// return
	if nil == left && right == nil {
		return true
	}
	if nil == left || right == nil {
		return false
	}
	if left.Val != right.Val {
		return false
	}
	return _isSymmetric(left.Left, right.Right) && _isSymmetric(left.Right, right.Left)// 判断对称节点
}

