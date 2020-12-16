package tree

/*
	根据一棵树的前序遍历与中序遍历构造二叉树。

	注意:
	你可以假设树中没有重复的元素。

	例如，给出

	前序遍历 preorder = [3,9,20,15,7]
	中序遍历 inorder = [9,3,15,20,7]
	返回如下的二叉树：

		3
	   / \
	  9  20
		/  \
	   15   7
*/
/*
	遍历preorder,取每一个值n,
	以n为root对inorder的左右元素再进行递归构建
	得到根节点在inorder中的定位索引
	算出左子节点数量
	算出右子节点数量
	递归
*/
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	root := &TreeNode{
		Val: preorder[0],
	}
	rootIndex := getIndex(inorder, root.Val)
	// build left
	root.Left = buildTree(preorder[1:len(inorder[:rootIndex])+1],inorder[:rootIndex+1])
	// build right
	root.Right = buildTree(preorder[len(inorder[:rootIndex])+1:], inorder[rootIndex+1:])
	return root
}



func getIndex(inorder []int, val int) int {
	for i, v := range inorder {
		if v == val {
			return i
		}
	}
	return -1
}