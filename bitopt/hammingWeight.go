package bitopt

/*
请实现一个函数，输入一个整数，输出该数二进制表示中 1 的个数。
例如，把 9 表示成二进制是 1001，有 2 位是 1。因此，如果输入 9，则该函数输出 2。
*/
//解法1 循环判断n的最后一位是否为1,直到n为0
func hammingWeight2(num uint32) int {
	var res = 0
	for num > 0 {
		if num & 1 == 1 {
			res++
		}
		num>>=1
	}
	return res
}
func hammingWeight(num uint32) int {
	var res = 0
	for num > 0 {
		res++
		num &= num-1
	}
	return res
}
