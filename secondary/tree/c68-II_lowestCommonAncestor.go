package tree

/*
	给定一个二叉树, 找到该树中两个指定节点的最近公共祖先。
中最近公共祖先的定义为：“对于有根树 T 的两个结点 p、q，最近公共祖先表示为一个结点 x，满足 x 是 p、q 的祖先且 x 的深度尽可能大（一个节点也可以是它自己的祖先）。”

例如，给定如下二叉树:  root = [3,5,1,6,2,0,8,null,null,7,4]
示例 1:

输入: root = [3,5,1,6,2,0,8,null,null,7,4], p = 5, q = 1
输出: 3
解释: 节点 5 和节点 1 的最近公共祖先是节点 3。
示例 2:

输入: root = [3,5,1,6,2,0,8,null,null,7,4], p = 5, q = 4
输出: 5
解释: 节点 5 和节点 4 的最近公共祖先是节点 5。因为根据定义最近公共祖先节点可以为节点本身。
 

说明:

所有节点的值都是唯一的。
p、q 为不同节点且均存在于给定的二叉树中。

*/


func LowestCommonAncestorII(root *TreeNode, p, q *TreeNode) *TreeNode {
	cur := root
	return _lowestCommonAncestorII(cur, p,q)
}
/*
	对于当前节点cur,如果p,q分别在其左子树和右子树中，则cur就是p,q的公共祖先
	采用递归的方式，如果cur与p或者q中的任何一个节点相等，则可以直接返回，因为此时存在以下情况(假设p==cur)
	q在cur的子节点中，则公共节点为cur
	q不在cur的子节点中，则此时根据题目条件，应该在以cur的父节点或者更上层节点为根的另一个子节点的子树中，则公共节点应该是cur的父节点或者更上层的节点
	所以可以根据以上逻辑检查cur的左右子节点是否包含p或者q，然后进行处理
*/
func _lowestCommonAncestorII(cur *TreeNode, p *TreeNode, q *TreeNode) *TreeNode {
	if nil == cur {
		return nil
	}
	if cur.Val == p.Val || cur.Val == q.Val {
		return cur
	}
	left := _lowestCommonAncestorII(cur.Left, p, q)
	right := _lowestCommonAncestorII(cur.Right, p, q)
	if nil != left && nil != right {
		return cur
	}
	if nil != left {
		return left
	}
	if nil != right {
		return right
	}
	return nil
}
