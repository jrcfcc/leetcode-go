package nums

/*
实现函数double Power(double base, int exponent)，求base的exponent次方。不得使用库函数，同时不需要考虑大数问题。

快速幂法:
当 x = 0时：直接返回 0 （避免后续 x = 1 / x 操作报错）。
初始化 res = 1 ；
当 n < 0时：把问题转化至n≥0 的范围内，即执行 x = 1/x ，n = - n ；
循环计算：当 n = 0时跳出；
当 n & 1 = 1 时：将当前 x 乘入 res （即 res *= x ）；
执行 x = x^2（即 x *= x）；
执行 n 右移一位（即 n >>= 1）。
返回 res 。
*/
func myPow(x float64, n int) float64 {
	if x == 0 {
		return x
	}
	var res = float64(1)
	if n < 0 {
		n = -n
		x = 1 / x
	}
	for n > 0 {
		if n & 1 == 1 {
			res *= x
		}
		x*=x
		n = n>>1
	}
	return res
}
