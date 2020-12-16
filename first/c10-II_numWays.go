package first
/*
	一只青蛙一次可以跳上1级台阶，也可以跳上2级台阶。求该青蛙跳上一个 n 级的台阶总共有多少种跳法。

	答案需要取模 1e9+7（1000000007），如计算初始结果为：1000000008，请返回 1。

	示例 1：

	输入：n = 2
	输出：2
	示例 2：

	输入：n = 7
	输出：21
	示例 3：

	输入：n = 0
	输出：1
	提示：

	0 <= n <= 100

*/
/*
	根据题义，对于第n级阶梯，青蛙可以从n-1级跳上来，也可以从n-2级跳上来
	所以f(n) = f(n-1)+f(n-2)
*/
func numWays(n int) int {
	var arr = make([]int, n+1)
	return _numWays(n, arr)
}

func _numWays(n int, arr []int) int {
	// return
	if n <= 1 {
		return 1
	}
	if arr[n] != 0 {
		return arr[n]
	}
	arr[n-1] = _numWays(n-1, arr)
	arr[n-2] = _numWays(n-2, arr)
	arr[n] = arr[n-1] + arr[n-2]
	return arr[n] % 1000000007

}