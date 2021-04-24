package main

import (
	"fmt"
	"github.com/themoonstone/toOffer/secondary/list"
)

var data []int = []int{1}

func main() {
	l1 := &list.ListNode{
		Val:  1,
		Next: &list.ListNode{
			Val:  2,
			Next: &list.ListNode{
				Val:  3,
				Next: &list.ListNode{
					Val:  4,
					Next: &list.ListNode{
						Val:  5,
						Next: nil,
					},
				},
			},
		},
	}

	fmt.Println(list.GetKthFromEnd(l1,2))
}

//var nodes = &tree.TreeNode{
//	Val: -2,
//	//Left: &toOffer.TreeNode{
//	//	Val: 4,
//	//	Left: &toOffer.TreeNode{
//	//		Val: 11,
//	//		Right: &toOffer.TreeNode{
//	//			Val: 2,
//	//		},
//	//		Left: &toOffer.TreeNode{
//	//			Val: 7,
//	//		},
//	//	},
//	//},
//	Right: &tree.TreeNode{
//		Val: 3,
//		//Right: &toOffer.TreeNode{
//		//	Val: 4,
//		//	Right: &toOffer.TreeNode{
//		//		Val: 1,
//		//	},
//		//	Left: &toOffer.TreeNode{
//		//		Val: 5,
//		//	},
//		//},
//		//Left: &toOffer.TreeNode{
//		//	Val: 13,
//		//},
//	},
//}

func maxDepth() {
	//toOffer.MaxDepth(nodes)
	//fmt.Println(toOffer.FindContinuousSequenceWithWindow(9))
}
