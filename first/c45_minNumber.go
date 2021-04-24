package first

import (
	"fmt"
)

/*
	输入一个非负整数数组，把数组里所有数字拼接起来排成一个数，打印能拼接出的所有数字中最小的一个。

	示例 1:

	输入: [10,2]
	输出: "102"
	示例 2:

	输入: [3,30,34,5,9]
	输出: "3033459"
	 

	提示:

	0 < nums.length <= 100
	说明:

	输出结果可能非常大，所以你需要返回一个字符串而不是整数
	拼接起来的数字可能会有前导 0，最后结果不需要去掉前导 0
*/
/*
	根据题意，应该先对数组进行分割，让个位数为0的在左边，不为0的在右边，然后分别排序
	对右边的数组
	按最高位进行排序
*/
func MinNumber(nums []int) string {
	return minNumber(nums)
}

// data := []int{20,1, 20,30}
func minNumber(nums []int) string {
	quickSort45(nums, 0, len(nums)-1)
	s := ""
	for i := 0; i <= len(nums)-1; i++ {
		s += fmt.Sprintf("%d", nums[i])
	}
	return s
}

func getData(a, b int) int {
	if fmt.Sprintf("%d%d",a,b) < fmt.Sprintf("%d%d",b,a) {
		return -1
	}
	return 1
}

func quickSort45(nums []int, l, r int) {
	if l >= r {
		return
	}
	p := partition45(nums, l, r)
	quickSort45(nums, l, p-1)
	quickSort45(nums, p+1, r)
}

func partition45(nums []int, l int, r int) int {
	v := nums[l]
	j := l
	for i := l + 1; i <= r; i++ {
		if getData(nums[i],v) < 0 {
			nums[i], nums[j+1] = nums[j+1], nums[i]
			j++
		}
	}
	nums[j], nums[l] = nums[l], nums[j]
	return j
}
