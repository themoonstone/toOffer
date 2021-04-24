package list
/*
	输入两个递增排序的链表，合并这两个链表并使新链表中的节点仍然是递增排序的。

	示例1：

	输入：1->2->4, 1->3->4
	输出：1->1->2->3->4->4
	限制：

	0 <= 链表长度 <= 1000
*/
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
/*
	双指针
*/
// n1 1->2->4
// n2 1->3->4
func MergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	return mergeTwoLists(l1, l2)
}
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	cur := new(ListNode) // 创建一个伪头节点
	lr := cur
	n1 := l1
	n2 := l2
	for  {
		if nil == n1 && nil == n2 {
			return lr.Next
		}
		if nil == n1 {
			cur.Next = n2
			cur = cur.Next
			n2 = n2.Next
		} else if nil == n2 {
			cur.Next = n1
			cur = cur.Next
			n1 = n1.Next
		}else if n1.Val < n2.Val {
			cur.Next = n1
			cur = cur.Next
			n1 = n1.Next
		} else {
			cur.Next = n2
			cur = cur.Next
			n2 = n2.Next
		}
	}
}

// recursive
func mergeTwoListsRecur(l1 *ListNode, l2 *ListNode) *ListNode{
	if nil == l1 && nil == l2 {
		return nil
	}
	if nil == l1 {
		return l2
	}
	if nil == l2 {
		return l1
	}
	if l1.Val < l2.Val {
		l1.Next = mergeTwoListsRecur(l1.Next, l2)
		return l1
	} else {
		l2.Next = mergeTwoListsRecur(l1, l2.Next)
		return l2
	}
}