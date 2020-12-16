package tree
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}


func _max(l int, r int) int {
	if l > r {
		return l
	}
	return r
}