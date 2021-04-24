package toOffer

import "math"

/*
	给定一个整数序列数组，确定它是否是严格递增的，或者通过删除一个元素使之成为严格递增序列。
	注：如果 a0 < a1 < … an 那么 a0, a1, …, an 也被认为是严格递增的序列。只包含一个元素的序列也被认为是严格递增的。

	示例
	输入 [1,3,2,1]，输出应该是 false。
	输入 [1,3,2]，输出应该是 true。你可以移除数组中的 3，得到序列 [1,2]；你也可以移除 2，得到序列 [1,3]；[1,2] 和 [1,3] 都是严格递增序列。
	输入输出
	[执行时间限制] 0.5 秒（cpp）
	[输入] 整数数组。约束：2 ≤ 序列长度 ≤ 105，-105 ≤ 序列 [i] ≤ 105。
	[输出] 布尔值
	如果从数组中删除一个元素可以使序列是严格递增的，则返回 true，否则返回 false。
*/
/*
	暴力法，直接遍历判断，遇到不符合条件的元素标记下标，然后删除之后重新遍历，满足条件即为真
*/
func almostIncreasingSequenceBOD(t1 []int)  bool {
	if len(t1) == 1 {
		return true
	}
	index := math.MaxInt32
	for i := 1; i < len(t1); i++ {
		if t1[i]<=t1[i-1] {
			index = i-1
		}
	}
	if index == math.MaxInt32 {
		return true
	}
	t1 = append(t1[:index],t1[index+1:]...)
	for i := 1; i < len(t1); i++ {
		if t1[i]<=t1[i-1] {
			return false
		}
	}
	return true
}


func min(a,b int) int {
	if a < b {
		return a
	}
	return b
}
func buMergeSort(nums []int) int {
	step := 1
	n := len(nums)
	count := 0
	for step<=n-1 {
		for i := 0; i < n; i+=2*step {
			count += merge(nums, i,(i+i+2*step-1)/2, min(i+2*step-1, n-1))
		}
		step*=2
	}

	return count
}

func merge(list []int, l int, mid int, r int) int {
	aux := make([]int, r-l+1)
	count := 0
	for i := l; i <= r; i++ {
		aux[i-l] = list[i]
	}
	i := l       // 指向左边有序的数组第一个元素
	j := mid + 1 // 指向右边有序数组的第一个元素
	for k := l; k <= r; k++ {
		// list[k]表示位置k merge(list, l, mid, r)当前要放的元素
		if i > mid {
			list[k] = aux[j-l]
			j++
		} else if j > r {
			// 注意，这里不能追加count,满足该条件说明右边所有的元素都已经比较完了，而在比较过程中，count已经完成了追加
			list[k] = aux[i-l]
			i++
		} else if aux[i-l] >= aux[j-l] {
			list[k] = aux[j-l]
			j++
			count+=(mid-i+1) // i,j都处于已经排好序的数组，lis[i]>list[j],则i后面的元素都大于list[j]
		} else {
			list[k] = aux[i-l]
			i++
		}
	}
	return count
}