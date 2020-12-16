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
	祖先的定义： 若节点 p 在节点 root 的左（右）子树中，或 p = root ，则称 root 是 p 的祖先。
	最近公共祖先的定义： 设节点 root 为节点 p, q的某公共祖先，若其左子节点 root.left 和右子节点 root.right 都不是 p,q的公共祖先，则称 root 是 “最近的公共祖先” 。
	根据以上定义，若 root 是 p, q 的 最近公共祖先 ，则只可能为以下情况之一：
	p 和 q 在 root 的子树中，且分列 root 的 异侧（即分别在左、右子树中）；
	p = root ，且 q 在 root 的左或右子树中；
	q = root ，且 p 在 root 的左或右子树中；
*/

func lowestCommonAncestorII(root *TreeNode, p, q *TreeNode) *TreeNode {
	cur := root
	return _lowestCommonAncestorII(cur, p,q)
}

// 递归后序遍历的思路
func _lowestCommonAncestorII(cur *TreeNode, p, q *TreeNode) *TreeNode {
	if cur == nil {
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
	if nil == left  && nil == right {
		return nil
	}
	if nil != left && nil == right {
		return left
	}
	if nil == left && nil != right {
		return left
	}
	return nil
}

/*
	遍历
	退出条件:cur=nil。
	然后循环判断p,q与cur的大小。
	如果p,q在root的左边，则cur=cur.left。
	如果p,q在root的右边，则cur=cur.right。
*/

func lowestCommonAncestor1II(root *TreeNode, p, q *TreeNode) *TreeNode {
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