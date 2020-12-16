package tree

import "fmt"

/*
	前序遍历 preorder = [3,9,20,15,7]
	中序遍历 inorder = [9,3,15,20,7]
*/
func buildTree(preorder []int, inorder []int) *TreeNode {
	if 0 == len(preorder) {
		return nil
	}
	root := &TreeNode{}
	rootValue := preorder[0]
	root.Val = rootValue
	index := getIndex(inorder, rootValue)
	fmt.Println("rootValue:", rootValue, "index:", index, "preorder:", preorder, "inorder:", inorder)
	root.Left = buildTree(preorder[1:len(inorder[:index])+1], inorder[:index])
	root.Right = buildTree(preorder[len(inorder[:index])+1:], inorder[index+1:])
	return root
}

func getIndex(inorder []int, value int) int {
	for index, v := range inorder {
		if v == value {
			return index
		}
	}
	return -1
}

func BuildTree(preorder []int, inorder []int) *TreeNode {
	return buildTree(preorder, inorder)
}
