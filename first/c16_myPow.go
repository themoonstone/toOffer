package first

/*
	实现函数double Power(double base, int exponent)，求base的exponent次方。不得使用库函数，同时不需要考虑大数问题。
 
	示例 1:

	输入: 2.00000, 10
	输出: 1024.00000
	示例 2:

	输入: 2.10000, 3
	输出: 9.26100
	示例 3:

	输入: 2.00000, -2
	输出: 0.25000
	解释: 2-2 = 1/22 = 1/4 = 0.25
	 

	说明:

	-100.0 < x < 100.0
	n 是 32 位有符号整数，其数值范围是 [−231, 231 − 1] 。

*/

func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	if n < 0 {
		x = 1 / x
		n = -1 * n
	}
	if n%2 == 0 {
		half := myPow(x, n/2)
		return half * half
	} else {
		// todo 为什么要分奇偶? 因为n是整数，如果n是奇数的话，n/2会被截取
		half := myPow(x, n/2)
		return half * half * x
	}

}

func MyPow(x float64, n int) float64 {
	//return myPow(x, n)
	return myPowIter(x, n)
}

// todo 逻辑没有理解
func myPowIter(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	result := 1.0
	if n < 0 {
		x = 1 / x
		n = -1 * n
	}
	for i := n; i != 0; i /= 2 {
		if i % 2 != 0 {
			//i是奇数
			result *= x
		}
		x *= x
	}
	return result
}
//
//class Solution {
//	public double myPow(double x, int n) {
//		double result = 1.0;
//		for (int i = n; i != 0; i /= 2, x *= x) {
//			if (i % 2 != 0) {
//				//i是奇数
//				result *= x;
//			}
//		}
//		return n < 0 ? 1.0 / result : result;
//	}
//}


// todo 内存溢出
//func myPow(x float64, n int) float64 {
//	if n == 0 {
//		return 1
//	}
//	if n > 0 {
//		return x * myPow(x, n-1)
//	} else {
//		return 1 / x * myPow(x, n+1)
//	}
//}
