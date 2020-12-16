package first

/*
	一个长度为n-1的递增排序数组中的所有数字都是唯一的，并且每个数字都在范围0～n-1之内。在范围0～n-1内的n个数字中有且只有一个数字不在该数组中，请找出这个数字。

	示例 1:

	输入: [0,1,3]
	输出: 2
	示例 2:

	输入: [0,1,2,3,4,5,6,7,9]
	输出: 8
	 

	限制：

	1 <= 数组长度 <= 10000

*/
func missingNumber(nums []int) int {
	for i, v := range nums {
		if i != v {
			return i
		}
	}
	return len(nums)
}

// 二分查找
/*
	判断中位数mid的值与nums[mid]是否相等，如果相等，则说明缺失的值在右边，如果不相等，则当前有可能是第一个缺失的值，也可能不是
*/
func missingNumber2(nums []int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		mid := (l + r) / 2
		if mid == nums[mid] {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return l
}
