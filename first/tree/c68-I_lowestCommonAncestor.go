package tree

/*
	给定一个二叉搜索树, 找到该树中两个指定节点的最近公共祖先。
	中最近公共祖先的定义为：“对于有根树 T 的两个结点 p、q，最近公共祖先表示为一个结点 x，满足 x 是 p、q 的祖先且 x 的深度尽可能大（一个节点也可以是它自己的祖先）。”

	例如，给定如下二叉搜索树:  root = [6,2,8,0,4,7,9,null,null,3,5]

	示例 1:

	输入: root = [6,2,8,0,4,7,9,null,null,3,5], p = 2, q = 8
	输出: 6
	解释: 节点 2 和节点 8 的最近公共祖先是 6。
	示例 2:

	输入: root = [6,2,8,0,4,7,9,null,null,3,5], p = 2, q = 4
	输出: 2
	解释: 节点 2 和节点 4 的最近公共祖先是 2, 因为根据定义最近公共祖先节点可以为节点本身。
	 

	说明:

	所有节点的值都是唯一的。
	p、q 为不同节点且均存在于给定的二叉搜索树中。
*/
/*
	递归:分条件判断, 如果p,q在root的两边，则其公共祖先就是root。
	如果p,q在root的左边，则递归左子节点判断。
	如果p,q在root的右边，则递归右子节点判断。
*/
func lowestCommonAncestor(root *TreeNode, p, q *TreeNode) TreeNode {
	cur := root
	return _lowestCommonAncestor(cur, p,q)
}

func _lowestCommonAncestor(cur *TreeNode, p, q *TreeNode) TreeNode {
	if p.Val<cur.Val && q.Val < cur.Val {
		return _lowestCommonAncestor(cur.Left, p ,q)
	} else if p.Val>cur.Val && q.Val > cur.Val {
		return _lowestCommonAncestor(cur.Right, p, q)
	}
	return *cur
}

/*
	遍历
	退出条件:cur=nil。
	然后循环判断p,q与cur的大小。
	如果p,q在root的左边，则cur=cur.left。
	如果p,q在root的右边，则cur=cur.right。
*/

func lowestCommonAncestor1(root *TreeNode, p, q *TreeNode) *TreeNode {
	cur := root
	for nil != cur {
		if p.Val<cur.Val && q.Val < cur.Val {
			cur = cur.Left
		} else if p.Val>cur.Val && q.Val > cur.Val {
			cur = cur.Right
		} else {
			break
		}
	}
	return cur
}