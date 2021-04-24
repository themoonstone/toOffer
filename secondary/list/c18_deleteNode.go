package list

func deleteNode(head *ListNode, val int) *ListNode {
	return recursiveDeleteNode(head, val)
}

// recursive
func recursiveDeleteNode(node *ListNode, val int) *ListNode {
	if nil == node {
		return nil
	}
	cur := recursiveDeleteNode(node.Next,val)
	if node.Val == val {
		return cur
	}
	node.Next = cur
	return node

}

// two point
func twopointDeleteNode(node *ListNode, val int) *ListNode {
	dummyHead := new(ListNode)
	dummyHead.Next = node
	prev := dummyHead
	for node != nil {
		if node.Val == val {
			prev.Next = node.Next
			node = nil
			return dummyHead.Next
		}
		prev = node
		node = node.Next
	}
	return nil
}