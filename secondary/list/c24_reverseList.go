package list

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

// recursive
func recursiveReverseList(cur *ListNode) *ListNode {
	if nil == cur || nil == cur.Next {
		return cur
	}
	res := recursiveReverseList(cur.Next) // 获取反转后的头节点
	cur.Next.Next = cur
	cur.Next = nil
	return res // 直接返回该头节点
}
// 双指针
/*
	给定一个pre保存前节点
	遍历当前节点
	让当前节点的next指向前节点
	更新前节点
*/
func reverseListTP(cur *ListNode) *ListNode {
	var pre *ListNode
	for nil != cur {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}