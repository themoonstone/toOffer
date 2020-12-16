package tree

func enqueue(queue []*TreeNode, node *TreeNode) []*TreeNode{
	return append(queue, node)
}
func enqueueFromHead(queue []*TreeNode, node *TreeNode) []*TreeNode {
	return append([]*TreeNode{node},queue[:]...)
}
func dequeue(queue []*TreeNode) ([]*TreeNode, *TreeNode) {
	if len(queue) == 1 {
		return nil, queue[0]
	}
	if len(queue) == 0 {
		return nil, nil
	}
	node := queue[0]
	queue = queue[1:len(queue)]
	return queue, node
}


func dqueue(queue []*TreeNode) ([]*TreeNode, *TreeNode) {
	if len(queue) == 0 {
		return nil, nil
	}
	return queue[1:], queue[0]
}
