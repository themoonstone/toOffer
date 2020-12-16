package first

import "fmt"

func ThirdMax(nums []int) int {
	seq := 3
	if len(nums) < 3 {
		seq = len(nums) - 1
	}
	return kMax(nums, seq-1)
}
func thirdMax(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	// 去重
	m_arr := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		m_arr[nums[i]] = i
	}
	data := make([]int, 0)
	for key, _ := range m_arr {
		data = append(data, key)
	}
	QuickSort(data) // 排序

	if len(data) < 3 {
		return data[len(data)-1]
	}

	return data[len(data)-3]
}

func kMax(nums []int, k int) int {
	return quickSortDesc(nums, 0, len(nums)-1, k)
}

func quickSortDesc(nums []int, l, r, k int) int {
	if l >= r {
		return nums[l]
	}
	p := descPartition(nums, l, r)
	if p == k {
		return nums[p]
	} else if p > k {
		return quickSortDesc(nums, l, p-1, k)
	} else {
		return quickSortDesc(nums, p+1, r, k)
	}
}

func descPartition(nums []int, l int, r int) int {
	v := nums[l]
	j := l
	// [l+1...j]>v, [j+1...i)<=v,i代表当前正在访问的元素
	for i := l + 1; i <= r; i++ {
		if nums[i] > v {
			nums[i], nums[j+1] = nums[j+1], nums[i]
			j++
		}
	}
	nums[j], nums[l] = nums[l], nums[j]
	return j
}

func FindKthLargest(nums []int, k int) int {
	// todo 注意:第k大是数字，换到数组下标需要减1
	n := quickSort(nums, k-1)
	fmt.Println(nums)
	return n
}

func quickSort(nums []int, k int) int {
	return _quickSort(nums, 0, len(nums)-1, k)
}

// 	data := []int{3,2,1,2}
func _quickSort(nums []int, l int, r int, k int) int {
	if l == r {
		return nums[l]
	}
	p := onePartition(nums, l, r)
	if p == k {
		return nums[p]
	} else if p > k {
		return _quickSort(nums, l, p-1, k)
	} else {
		return _quickSort(nums, p+1, r, k)
	}
}

func QuickSort(nums []int) {
	_quickSort1(nums, 0, len(nums)-1)

}

func _quickSort1(nums []int, l int, r int) {
	if l >= r {
		return
	}
	p := _partition(nums, l, r)
	//lt, gt := _threeWayPartition(nums, l, r)

	_quickSort1(nums, l, p-1)
	_quickSort1(nums, p+1, r)
}

func partition(nums []int, l, r int) int {
	v := nums[l]
	i, j := l+1, r
	for i <= j {
		if nums[i] <= v {
			i++
		} else {
			nums[i], nums[j] = nums[j], nums[i]
			j--
		}
	}
	nums[l], nums[i-1] = nums[i-1], nums[l]
	return i - 1
}

func onePartition(nums []int, l, r int) int {
	v := nums[l]
	j := l // [l+1...j]<v, [j+1...i-1]>=v,i代表当前访问的元素
	for i := l + 1; i <= r; i++ {
		if nums[i] < v {
			nums[j+1], nums[i] = nums[i], nums[j+1]
			j++
		}
	}
	nums[l], nums[j] = nums[j], nums[l]
	return j
}

// 倒序快排
func _partition(nums []int, l int, r int) int {
	j := l
	v := nums[l]
	for i := l + 1; i <= r; i++ {
		if nums[i] > v {
			nums[i], nums[j+1] = nums[j+1], nums[i]
			j++
		}
	}
	nums[l], nums[j] = nums[j], nums[l]
	return j
}

// 双路快排
// 定义i, j 为当前两边访问的元素，则对于整个数组来说，[l+1...i)范围的元素都应该小于等于v,(j...r]范围的元素都应该大于等于v
func _twoWayPartition(nums []int, l, r int) int {
	j := r
	v := nums[l]
	i := l + 1
	for {
		for i <= r && nums[i] < v {
			i++
		}
		for j >= l+1 && nums[j] > v {
			j--
		}
		if i > j {
			nums[j], nums[l] = nums[l], nums[j]
			return j
		}
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
}
func _twoWayPartition2(nums []int, l, r int) int {
	v := nums[l]
	i, j := l+1, r
	for {
		if i > j {
			break
		}
		for nums[i] < v && i <= r {
			i++
		}
		for nums[j] > v && j >= l {
			j--
		}
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
	nums[l], nums[i-1] = nums[i-1], nums[l]
	return i
}

// 三路快排
/*
	data := []int{3,2,1,3,45,6,23,3}

	[l+1...lt]<v
	[lt+1...i)=v
	[gt...r]>v
*/

func _threeWayPartition(nums []int, l, r int) (int, int) {
	//lt := l     // [l+1...lt]<v
	lt := l
	gt := r + 1 // [gt...r]>v
	v := nums[l]
	for i := l + 1; i < gt; {
		if nums[i] > v {
			nums[i], nums[gt-1] = nums[gt-1], nums[i]
			gt--
		} else if nums[i] < v {
			nums[i], nums[lt+1] = nums[lt+1], nums[i]
			i++
			lt++
		} else {
			i++
		}
	}
	nums[l], nums[lt] = nums[lt], nums[l]
	lt--
	return lt, gt
}
