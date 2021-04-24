package secondary

/*
	在数组中的两个数字，如果前面一个数字大于后面的数字，则这两个数字组成一个逆序对。输入一个数组，求出这个数组中的逆序对的总数。
	示例 1:

	输入: [7,5,6,4]
	输出: 5
	限制：

	0 <= 数组长度 <= 50000
*/

func reversePairs(nums []int) int {
	return buMergeSort(nums)
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
		} else if aux[i-l] > aux[j-l] {
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