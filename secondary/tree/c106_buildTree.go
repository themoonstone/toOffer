package tree

import "fmt"

/*
	根据一棵树的中序遍历与后序遍历构造二叉树。

	注意:
	你可以假设树中没有重复的元素。

	例如，给出

	中序遍历 inorder = [9,3,15,20,7]
	后序遍历 postorder = [9,15,7,20,3]
	返回如下的二叉树：

		3
	   / \
	  9  20
		/  \
	   15   7

*/

func buildTreePost(inorder []int, postorder []int) *TreeNode {
	if len(postorder) == 0 {
		return nil
	}
	root := &TreeNode{}
	rootValue := postorder[len(postorder)-1]
	root.Val = rootValue
	index := getIndex(inorder, rootValue)
	fmt.Println("index:", index, "inorder:", inorder, "postorder:", postorder, "rootValue:", rootValue)
	root.Left = buildTreePost(inorder[:index], postorder[:len(inorder[:index])])
	root.Right = buildTreePost(inorder[index+1:], postorder[len(inorder[:index]):len(postorder)-1])
	return root
}

func BuildTreePost(inorder []int, postorder []int) *TreeNode {
	return buildTreePost(inorder, postorder)
}
