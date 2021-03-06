package secondary

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
/*
	根据题义，可以将数组分成两块，缺失元素前的数组nums[i]=i,从缺失开始的数组nums[i]>i
*/
func missingNumber(nums []int) int{
	i, j := 0, len(nums) - 1
	for i <= j {
		mid := (i+j)/2
		if nums[mid] == mid {
			i = mid + 1
		} else {
			j = mid - 1
		}
	}
	return i
}