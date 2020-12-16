package tree

/*
	输入两棵二叉树A和B，判断B是不是A的子结构。(约定空树不是任意一个树的子结构)

	B是A的子结构， 即 A中有出现和B相同的结构和节点值。

	例如:
	给定的树 A:

	     3
	    / \
	   4   5
	  / \
	 1   2
	给定的树 B：

	   4 
	  /
	 1
	返回 true，因为 B 与 A 的一个子树拥有相同的结构和节点值。

	示例 1：

	输入：A = [1,2,3], B = [3,1]
	输出：false
	示例 2：

	输入：A = [3,4,5,1,2], B = [4,1]
	输出：true
	限制：

	0 <= 节点个数 <= 10000

*/

/*
	前序遍历
	遍历A的每一个节点，判断以当前节点为根的树是否包含了B
*/
func isSubStructure(A *TreeNode, B *TreeNode) bool {
	if nil == B || nil == A {
		return false
	}
	return recur(A, B) || isSubStructure(A.Left, B) || isSubStructure(A.Right, B)
}

func recur(a *TreeNode, b *TreeNode) bool {
	if nil == b {
		return true
	}
	if nil == a || a.Val != b.Val {
		return false
	}
	return recur(a.Left, b.Left) && recur(a.Right, b.Right)
}
