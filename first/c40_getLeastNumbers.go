package first

import "fmt"

/*
	输入整数数组 arr ，找出其中最小的 k 个数。例如，输入4、5、1、6、2、7、3、8这8个数字，则最小的4个数字是1、2、3、4。

	示例 1：

	输入：arr = [3,2,1], k = 2
	输出：[1,2] 或者 [2,1]
	示例 2：

	输入：arr = [0,1,2,1], k = 1
	输出：[0]
	 

	限制：

	0 <= k <= arr.length <= 10000
	0 <= arr[i] <= 10000

*/

/*
	通过快排定位k的位置，根据快排的定义思想，找到k之后，说明k左边的数都比arr[k]小，所以，最小的k个数为arr[0...k)
*/
func quickSort40(arr []int, k int) []int {
	 _quickSort40(arr, 0, len(arr)-1, k-1)
	 return arr[:k]
}

func _quickSort40(arr []int, l int, r int, k int) int {
	if l >= r {
		return l
	}
	p := _partition40(arr, l, r, k)
	if p == k {
		return p
	} else if p > k {
		return _quickSort40(arr, l, p-1, k)
	} else {
		return _quickSort40(arr, p+1, r, k)
	}
}

func _partition40(arr []int, l int, r int, k int) int {
	v := arr[l]
	j := l                        // 大于v和小于等于v的分解点
	for i := l + 1; i <= r; i++ { // i表示当前访问的元素下标
		if arr[i] <= v {
			arr[j+1], arr[i] = arr[i], arr[j+1]
			j++
		}
	}
	arr[l], arr[j] = arr[j], arr[l]
	return j
}

/*
	实现一个最小堆
*/
func GetLeastNumber(arr []int, k int) []int {
	return getLeastNumbers(arr, k)
}
func getLeastNumbers(arr []int, k int) []int {
	result := make([]int, 0)
	heap := Construct(Heap{})
	for i := 0; i < len(arr); i++ {
		heap = Insert(heap, arr[i])
	}
	fmt.Println("heap:", heap.Datas)

	for i := 0; i < k; i++ {
		var data int
		heap, data = Remove(heap)
		result = append(result, data)
	}
	return result
}

type Heap struct {
	Datas []int
	size  int
}

func Construct(heap Heap) Heap {

	return Heap{
		Datas: make([]int, 1),
		size:  0,
	}
}

func IsEmpty(heap Heap) bool {
	return heap.size <= 0
}

func Count(heap Heap) int {
	return heap.size
}

func Insert(heap Heap, data int) Heap {
	heap.Datas = append(heap.Datas, data)
	heap.size++
	heap = shiftUp(heap, data)
	return heap
}

func shiftUp(heap Heap, data int) Heap {
	index := len(heap.Datas) - 1
	for index > 1 && heap.Datas[index] < heap.Datas[index/2] {
		heap.Datas[index], heap.Datas[index/2] = heap.Datas[index/2], heap.Datas[index]
		index /= 2
	}
	return heap
}

func Remove(heap Heap) (Heap, int) {
	if IsEmpty(heap) {
		return heap, 0
	}
	// 交换堆顶元素与最后一个元素
	heap.Datas[1], heap.Datas[len(heap.Datas)-1] = heap.Datas[len(heap.Datas)-1], heap.Datas[1]
	data := heap.Datas[len(heap.Datas)-1]
	// 删除最后一个元素
	heap.Datas = heap.Datas[:len(heap.Datas)-1]
	heap.size--
	// shiftDown
	shiftDown(heap, 1)
	return heap, data
}

func shiftDown(heap Heap, index int) {
	for index*2 <= len(heap.Datas)-1 {
		j := 2 * index
		if j+1 <= len(heap.Datas)-1 && heap.Datas[index*2] > heap.Datas[2*index+1] {
			j++ // 有右孩子且右孩子小于左孩子
		}
		if heap.Datas[index] <= heap.Datas[j] {
			break
		}
		heap.Datas[index], heap.Datas[j] = heap.Datas[j], heap.Datas[index]
		index = j
	}
}
