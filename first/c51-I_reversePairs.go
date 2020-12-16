package first

/*
	方法1：暴力破解
	外层从0开始，内层从i开始，判断每一个数组是否是逆序对，进行统计
	时间复杂度:O(n^2)
	todo leetcode提交中超出时间限制
*/
func reversePairs(nums []int) int {
	count := 0
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] > nums[j] {
				count++
			}
		}
	}
	return count
}

/*
	方法二：原始解法时间复杂度太高，可以采用归并排序
	对于一个已经归并好的数组，肯定是有序的，所以在进行归并过程中，对元素进行比较的时候，
	如果右边的元素l[j]小于左边的元素l[i]，则说明它比左边数组中的l[i]后面的元素都要小，
	因此到[j]处的逆序对数量就是len[l[j:r])的长度
*/
func ReversePairsMerge(nums []int) int {
	var count int
	 Sort(nums, 0, len(nums)-1, &count)
	return count
}

func Sort(nums []int, l, r int, count *int) {
	if l >= r {
		return
	}
	mid := (l + r) / 2

	Sort(nums, l, mid,count)
	Sort(nums, mid+1, r,count)
	merge(nums, l, mid, r,count)
}

func merge(nums []int, l int, mid int, r int,count*int) {
	i := l
	j := mid + 1
	arr := make([]int, r-l+1)
	for k := l; k <= r; k++ {
		arr[k-l] = nums[k]
	}
	// 对arr进行比较
	for k := l; k <= r; k++ {
		if i > mid {
			nums[k] = arr[j-l]
			j++
		} else if j > r {
			nums[k] = arr[i-l]
			i++
		} else if arr[i-l] <= arr[j-l] { // todo 注意：这个时候必需加上等于，因为如果将等于放到大于那一边判断会将两个相同的元素也判断为逆序对
			nums[k] = arr[i-l]
			i++
		} else {
			nums[k] = arr[j-l]
			j++
			*count += (mid - i + 1)
		}
	}
}
