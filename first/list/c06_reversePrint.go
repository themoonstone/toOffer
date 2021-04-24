package list
/*
	输入一个链表的头节点，从尾到头反过来返回每个节点的值（用数组返回）。

	示例 1：

	输入：head = [1,3,2]
	输出：[2,3,1]
	 

	限制：

	0 <= 链表长度 <= 10000
*/

// 先反转链表，再遍历打印
func reversePrint(head *ListNode) []int {
	var node *ListNode
	node = reversePrintList(head)
	result := make([]int, 0)
	for nil != node {
		result = append(result, node.Val)
		node = node.Next
	}
	return result
}

func reversePrintList(node *ListNode) *ListNode {
	if nil == node || nil == node.Next {
		return node
	}
	ret := reversePrintList(node.Next)
	node.Next.Next = node
	node.Next = nil
	return ret
}

// 用栈解决
func reversePrintStack(node *ListNode) []int {
	stack := make([]int, 0)
	result := make([]int, 0)

	for nil != node {
		stack = append(stack, node.Val)
		node = node.Next
	}
	for i := len(stack) - 1; i >= 0;i -- {
		result = append(result, stack[i])
	}
	return result
}