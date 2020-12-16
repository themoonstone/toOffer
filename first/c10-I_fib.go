package first

func Fib2(n int) int {
	if n <= 1 {
		return n
	}
	return Fib2(n - 1) + Fib(n - 2)
}

func Fib(n int) int {
	if n <= 1 {
		return n
	}
	var arr = make([]int, n+1)
	return _fib(n, arr)
}

func _fib(n int, arr []int) int {
	// return
	if n <= 1 {
		return n
	}
	if arr[n] != 0 {
		return arr[n]
	}
	arr[n-1] = _fib(n-1, arr)
	arr[n-2] = _fib(n-2, arr)
	arr[n] = arr[n-1] + arr[n-2]
	return arr[n]

}
