package list

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

/*
	递归:要反转链表，需要当前节点的下一个节点的next指向自己，但不能从一开始就这么做，因为这样会导致链表断掉
	所以可以采用递归的写法，从最后一个开始向前走
*/
func reverseList(head *ListNode) *ListNode {
	return _reverseList(head)
}

func _reverseList(node *ListNode) *ListNode {
	// return
	if node == nil || node.Next == nil {
		// 找到最后一个节点或者链表本身是一个空链表
		return node
	}
	// current logic
	// 获取头节点
	ret := _reverseList(node.Next)
	// 让当前节点的下一个节点的next指向当前节点
	node.Next.Next = node
	// 将当前节点的next置空，防止循环链表
	node.Next = nil
	return ret
}

/*
	双指针
*/
func _reverseListPoint(node *ListNode) *ListNode{
	var pre *ListNode// 保存前节点指针
	for node != nil {
		next := node.Next
		node.Next = pre // 让当前节点指向前节点
		pre = node // 移动前节点指针
		node = next // 移动当前节点指针
	}
	return pre
}
