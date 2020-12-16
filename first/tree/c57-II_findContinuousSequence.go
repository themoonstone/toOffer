package tree

// 1 暴力破解
func findContinuousSequence(target int) [][]int {
	result := make([][]int, 0)
	n := target/2 + 1
	for i := 1; i <= n; i++ {
		single := []int{i}
		res := i
		for j := i + 1; j <= n; j++ {
			res += j
			if res < target {
				single = append(single, j)
			} else if res == target && j-i >= 1 {
				single = append(single, j)
				result = append(result, single)
				break
			} else {
				break
			}
		}
	}
	return result
}

/*
	2 滑动窗口
	设定一个滑动窗口，包含了数组中的一部分元素，对其中的元素进行观察，如果和值sum>target，则左边界i右移
	如果sum<targe,则右边界j右移，通过这种方式可以得到所有和为target的情况
	时间复杂度表O(n)
*/
func FindContinuousSequenceWithWindow(target int) [][]int {
	result := make([][]int, 0)
	n := target/2 + 1
	j := 1
	sum := 0
	for i := 1; i <= n ; {

		if sum > target {
			sum -= i
			i++
		}
		if sum < target {
			sum += j
			j++
		}
		if sum == target{
			single := make([]int, 0)
			for k := i; k < j; k++ {
				single = append(single, k)
			}
			result = append(result, single)
			sum -= i
			i++
		}
	}
	return result
}
