package secondary

/*
	统计一个数字在排序数组中出现的次数。

	示例 1:

	输入: nums = [5,7,7,8,8,10], target = 8
	输出: 2
	示例 2:

	输入: nums = [5,7,7,8,8,10], target = 6
	输出: 0
	 

	限制：

	0 <= 数组长度 <= 50000
*/
/*
	通过二分查找找到与target相等的下标，然后分别向左和向右遍历，找到其左边界l和右边界r
*/
func search(nums []int, target int) int {
	l, r := 0, 0
	i, j := 0, len(nums)-1
	// 找到左边界
	for i <= j {
		mid := (i + j) / 2
		if nums[mid] >= target {
			j = mid - 1
		} else {
			i = mid + 1
		}
	}
	l = j
	// 找到右边界
	i ,j = 0, len(nums) - 1
	for i <= j {
		mid := (i+j)/2
		if nums[mid] <= target {
			i  = mid + 1
		} else {
			j = mid - 1
		}
	}
	r = i
	return r - l - 1
}
