package nums

/*
输入数字 n，按顺序打印出从 1 到最大的 n 位十进制数。比如输入 3，
则打印出 1、2、3 一直到最大的 3 位数 999。
说明：
用返回一个整数列表来代替打印
n 为正整数
*/
func printNumbers(n int) []int {
	if n <= 0 {
		return nil
	}
	var max = int(^uint(0) >> 1)
	var nums = int64(1)
	for n > 0 {
		nums = nums * 10
		n--
		if int64(max/10) < nums {
			nums = int64(max)
			break
		}
	}
	var l = int(nums-1)
	var res = make([]int,l)
	for i:=0;i<l;i++{
		res[i] = i+1
	}
	return res
}
