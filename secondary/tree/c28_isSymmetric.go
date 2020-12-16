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
	找对称节点，判断其值是否相等
*/
func isSymmetric(root *TreeNode) bool {
	if nil == root {
		return true
	}
	return _isSymmetric(root.Left, root.Right)
}

// 需要注意的是，这里的参数不是指根节点的左右子节点，指的是两个对称节点
func _isSymmetric(left *TreeNode, right *TreeNode) bool {
	if nil == left && nil == right {
		return true
	}
	if nil == left || nil == right {
		return false
	}
	if left.Val != right.Val {
		return false
	}
	return _isSymmetric(left.Left, right.Right) && _isSymmetric(left.Right, right.Left)
}
