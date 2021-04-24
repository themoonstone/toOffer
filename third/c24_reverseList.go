package third
/*
	定义一个函数，输入一个链表的头节点，反转该链表并输出反转后链表的头节点。

	示例:

	输入: 1->2->3->4->5->NULL
	输出: 5->4->3->2->1->NULL
	 

	限制：

	0 <= 节点个数 <= 5000

*/
func ReverseList(head *ListNode) *ListNode {
	return recursiveReverseList(head)
}

func recursiveReverseList(node *ListNode) *ListNode {
	if nil == node || nil == node.Next {
		return node
	}
	ret := recursiveReverseList(node.Next)
	node.Next.Next = node
	node.Next = nil
	return ret
}

// 双指针
/*
	遍历链表，以cur表示当前访问的节点,next表示cur的下一个节点,pre表示cur的前一个节点
	先保存next，然后断开cur与next的连接，让cur.next指向pre,实现当前cur的反转，然后更新pre
*/
func twoReverseList(cur *ListNode) *ListNode {
	var pre *ListNode
	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre= cur
		cur = next
	}
	return pre
}