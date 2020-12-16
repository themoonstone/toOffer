package first

import "sort"

/*
	找出数组中重复的数字。


	在一个长度为 n 的数组 nums 里的所有数字都在 0～n-1 的范围内。数组中某些数字是重复的，但不知道有几个数字重复了，也不知道每个数字重复了几次。请找出数组中任意一个重复的数字。

	示例 1：

	输入：
	[2, 3, 1, 0, 2, 5, 3]
	输出：2 或 3
	 

	限制：

	2 <= n <= 100000

*/
// 采用map存储数据
func findRepeatNumberMap(nums []int) int {
	mapNums := make(map[int]bool)
	for _, num := range  nums {
		if mapNums[num] {
			return num
		}
		mapNums[num] = true
	}
	return 0
}

func findRepeatNumberMapToPoint(nums []int) int {
	sort.Ints(nums)
	pre := nums[0]
	repeat := -1
	for i := 1 ; i < len(nums);i++ {
		if pre == nums[i] {
			repeat = pre
			break
		}
		pre = nums[i]
	}
	return repeat
}