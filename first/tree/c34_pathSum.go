package tree

/*
	输入一棵二叉树和一个整数，打印出二叉树中节点值的和为输入整数的所有路径。从树的根节点开始往下一直到叶节点所经过的节点形成一条路径。
	示例:
	给定如下二叉树，以及目标和 sum = 22，

				  5
				 / \
				4   8
			   /   / \
			  11  13  4
			 /  \    / \
			7    2  5   1
	返回:

	[
	   [5,4,11,2],
	   [5,8,4,5]
	]

*/

/*
	DFS+回塑
*/
var visitedNode  = make(map[*TreeNode]bool)
func PathSum(root *TreeNode, sum int) [][]int {
	var result = make([][]int, 0)
	var sumSlice  = make([]int,0 )
	var curSum int = 0
	_pathSum(root , sum, curSum, sumSlice, &result)
	return result
}

func _pathSum(node *TreeNode, sum int, curSum int,  sumSlice []int, result *[][]int) {
	if visitedNode[node] || nil == node{
		return
	}
	if curSum + node.Val == sum && node.Left == nil && node.Right == nil {
		// 最后一个节点必须是叶子节点才能形成一条路径
		newSlice := make([]int, 0)
		newSlice = append(newSlice,sumSlice[:]...)
		newSlice = append(newSlice, node.Val)
		*result = append(*result, newSlice)
		if visitedNode[node.Right] {
			visitedNode[node] = true
		}
		return
	} else {
		curSum = curSum + node.Val
		sumSlice = append(sumSlice, node.Val)
		_pathSum(node.Left, sum, curSum, sumSlice, result)
		_pathSum(node.Right, sum, curSum, sumSlice, result)
		return
	}
}

func getAbs()  {
	
}