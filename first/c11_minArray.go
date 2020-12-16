package first

/*
	把一个数组最开始的若干个元素搬到数组的末尾，我们称之为数组的旋转。输入一个递增排序的数组的一个旋转，输出旋转数组的最小元素。例如，数组 [3,4,5,1,2] 为 [1,2,3,4,5] 的一个旋转，该数组的最小值为1。  

	示例 1：
	输入：[3,4,5,1,2]
	输出：1

	示例 2：
	输入：[2,2,2,0,1]
	输出：0
*/
/*
	题解:对于一个递增排序数组的旋转数组来说，它应该由两个有序递增数组构成。此时对于旋转数组的最后一个元素x存在如下特点:
	1. 对数组中任何一个元素，存在y>x则说明y应该在最小值左侧，y<=x则说明y就是最小值或者在最小值右侧，所以可以采用二分查找

*/
func minArray0(numbers []int) int {
	low, hight := 0 , len(numbers) - 1
	for low < hight {
		mid := (low + hight) / 2
		if numbers[mid] > numbers[hight] {
			low = mid + 1
		} else if numbers[mid] < numbers[hight] {
			hight = mid
		} else {
			hight--
		}
	}
	return numbers[low]
}
func MinArray(numbers []int) int {
	return minArray(numbers)
}

/*
	先用快排将数据进行排序，然后取第一个元素
*/
func minArray(numbers []int) int {
	quickSortN(numbers)
	return numbers[0]
}

func quickSortN(nums []int) {
	_quickSortN(nums, 0, len(nums)-1)
}

func _quickSortN(nums []int, l int, r int) {
	if l >= r {
		return
	}
	k := _partitionN(nums, l, r) // k 表示k左边的元素都小于nums[k],k右边的元素都大于等于nums[k]
	_quickSortN(nums, l, k-1)
	_quickSortN(nums, k+1, r)
}

func _partitionN(nums []int, l int, r int) int {
	v := nums[l]
	i, j := l+1, l // i表示当前正在访问的元素下标，j表示大于v和小于v的分界线(此处为最后一个小于v的下标)

	for i <= r {
		if nums[i] < v {
			j++
			nums[i], nums[j] = nums[j], nums[i]
		}
		i++
	}
	nums[l], nums[j] = nums[j], nums[l]
	return j
}
