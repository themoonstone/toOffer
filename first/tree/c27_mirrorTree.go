package tree


// 要输出二叉树的镜像，其实就是左右子树对调，并一直递归
func mirrorTree(root *TreeNode) *TreeNode {
	// return
	if nil == root {
		return nil
	}
	// current logic
	root.Left, root.Right = root.Right, root.Left
	mirrorTree(root.Left)
	mirrorTree(root.Right)
	return root
}

/*
	广度优先:BFS+QUEUE
*/
func BFSMirrorTree(root *TreeNode) *TreeNode {
	if nil == root {
		return nil
	}
	queue := []*TreeNode{root}
	cur_node := &TreeNode{}
	for len(queue) != 0 {
		queue, cur_node = dqueue(queue)
		if nil != cur_node {
			cur_node.Left, cur_node.Right = cur_node.Right, cur_node.Left
			queue = append(queue,cur_node.Left,cur_node.Right)
		}
	}
	return root
}
