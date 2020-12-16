package first

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

var mNums = make(map[int]int)

func search(nums []int, target int) int {
	for _, v := range nums {
		mNums[v]++
	}
	return mNums[target]
}

func Search2(nums []int, target int) int {

	l, r := 0, len(nums)-1
	count := 0
	tIndex := _search(nums, l, r, target)
	if tIndex == -1 {
		return 0
	}
	for i := tIndex; i < len(nums); i++ {
		if nums[i] == nums[tIndex] {
			count++
		} else {
			break
		}
	}
	for j := tIndex-1; j >= 0; j-- {
		if nums[j] == nums[tIndex] {
			count++
		} else {
			break
		}
	}
	return count
}

func _search(nums []int, l int, r int, target int) int {
	mid := (l + r) / 2
	for l <= r{
		if nums[mid] > target {
			r = mid - 1
			mid = (l + r) / 2
		} else if nums[mid] < target {
			l = mid + 1
			mid = (l + r) / 2
		} else {
			return mid
		}
	}
	return -1

}

func _search2(nums []int, l int, r int, target int) int {
	mid := (l + r) / 2

	for l <= r{
		if nums[mid] > target {
			r = mid - 1
			mid = (l + r) / 2
		} else if nums[mid] < target {
			l = mid + 1
			mid = (l + r) / 2
		} else {
			return mid
		}
	}
	return -1

}

/*
	边界问题:确定要在一个什么样的范围区间中进行查找
	循环不变量的定义维护
	在[l...r]的闭区间中查找指定目标target
	在声明变量的时候要明确变量的定义，同时需要在程序执行过程中不断的维护这个变量的定义
*/
func binarySearch(nums []int, target int) int {
	l, r := 0, len(nums) - 1
	for l <= r { // 这里的范围定义也取决于区间，因为是在[l...r]的闭区间中查找，则在l==r的情况下，[l:r]仍是一个长度为1的有效数组，所以要包含等于
		mid := (l+r)/ 2
		if nums[mid] == target {
			return mid
		}
		if nums[mid] > target {
			r  = mid - 1
		} else {
			l = mid + 1
		}
	}
	return -1
}

/*
	在[l...r)前闭后开的区间中进行查找
	初始值 l = 0, r = n
	退出条件: 因为是在[l...r)的半闭半开区间中查找，则在l==r的情况下，[l:r)已经不是一个有效数组，所以不能再包含等于
*/
func binarySearch1(nums []int, target int) int {
	l, r := 0, len(nums)
	for l < r {
		mid := (l+r)/ 2
		if nums[mid] == target {
			return mid
		}
		if nums[mid] > target {
			r  = mid // 需要注意的是，此时因为右边是一个开区间，此时target应该是包含在[l...mid)这个半闭半开的区间中，
			// 而此时nums[mid]也不会包含在这个区间中，如果此处写成 r = mid-1,根据区间的范围，就会漏掉一个arr[mid-1]的元素
		} else {
			l = mid + 1
		}
	}
	return -1
}