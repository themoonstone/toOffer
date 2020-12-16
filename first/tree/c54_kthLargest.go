package tree

/*
	给定一棵二叉搜索树，请找出其中第k大的节点。

 

示例 1:

输入: root = [3,1,4,null,2], k = 1
   3
  / \
 1   4
  \
   2
输出: 4
示例 2:

输入: root = [5,3,6,2,4,null,null,1], k = 3
       5
      / \
     3   6
    / \
   2   4
  /
 1
输出: 4
 

限制：

1 ≤ k ≤ 二叉搜索树元素个数

*/
/*
	通过中序遍历将其存入一个数组中，则该数据一定是升序的，然后根据下标k的到倒数的元素即可
*/
func kthLargest(root *TreeNode, k int) int {
	array := make([]int, 0)
	inOrder(&array, root)
	return array[len(array) - k] // 第1大即最后一个元素，也就是array[len(array)-1]
}

func inOrder(array *[]int, root *TreeNode) {
	if nil == root {
		return
	}
	inOrder(array, root.Left)
	*array = append(*array, root.Val)
	inOrder(array, root.Right)
}