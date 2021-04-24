package list
/*
	输入一个链表的头节点，从尾到头反过来返回每个节点的值（用数组返回）。

	示例 1：

	输入：head = [1,3,2]
	输出：[2,3,1]
	 

	限制：

	0 <= 链表长度 <= 10000
*/

func reversePrint(head *ListNode) []int {
	head = recursiveReversePrint(head)
	var result []int = make([]int, 0)
	for nil != head {
		result = append(result, head.Val)
		head = head.Next
	}
	return result
}

func recursiveReversePrint(node *ListNode) *ListNode {
	if nil == node || nil == node.Next {
		return node
	}
	ret := recursiveReversePrint(node.Next)
	node.Next.Next = node
	node.Next = nil
	return ret
}
