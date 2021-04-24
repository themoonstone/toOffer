package list
/*
	给定单向链表的头指针和一个要删除的节点的值，定义一个函数删除该节点。

	返回删除后的链表的头节点。

	注意：此题对比原题有改动

	示例 1:

	输入: head = [4,5,1,9], val = 5
	输出: [4,1,9]
	解释: 给定你链表中值为 5 的第二个节点，那么在调用了你的函数之后，该链表应变为 4 -> 1 -> 9.
	示例 2:

	输入: head = [4,5,1,9], val = 1
	输出: [4,5,9]
	解释: 给定你链表中值为 1 的第三个节点，那么在调用了你的函数之后，该链表应变为 4 -> 5 -> 9.
*/

func deleteNode(head *ListNode, val int) *ListNode {
	return violentDelete(head,val)
}

// O(n)
func violentDelete(head *ListNode, val int) *ListNode {
	dumpNode:= new(ListNode) // 虚拟头节点
	dumpNode.Next = head
	var prev *ListNode = dumpNode
	for head != nil {
		if val==head.Val {
			prev.Next = head.Next
			head = nil
			return dumpNode.Next
		}
		prev = head
		head = head.Next
	}
	return nil
}

func recursiveDelete(head *ListNode, val int) *ListNode {
	if nil == head {
		return nil
	}
	resultNode := recursiveDelete(head.Next, val) // 存储的是头节点后面跟的那个链表中所有的值为val的节点删除后剩下的这个链表
	if head.Val == val {
		return resultNode
	} else {
		head.Next = resultNode
		return head
	}
}
