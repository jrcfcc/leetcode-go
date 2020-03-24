package strchar

import "math"

/*
请你来实现一个 atoi 函数，使其能将字符串转换成整数。
首先，该函数会根据需要丢弃无用的开头空格字符，直到寻找到第一个非空格的字符为止。
当我们寻找到的第一个非空字符为正或者负号时，则将该符号与之后面尽可能多的连续数字组合起来，
作为该整数的正负号；假如第一个非空字符是数字，则直接将其与之后连续的数字字符组合起来，形成整数。
该字符串除了有效的整数部分之后也可能会存在多余的字符，这些字符可以被忽略，它们对于函数不应该造成影响。
注意：假如该字符串中的第一个非空格字符不是一个有效整数字符、字符串为空或字符串仅包含空白字符时，
则你的函数不需要进行转换。
在任何情况下，若函数不能进行有效的转换时，请返回 0。
说明：
假设我们的环境只能存储 32 位大小的有符号整数，那么其数值范围为 [−231,  231 − 1]。
如果数值超过这个范围，请返回  INT_MAX (231 − 1) 或 INT_MIN (−231)
*/

/*
对数a取反，结果为-（a+1）
*/
const Max = int32(^uint32(0) >> 1)
const Min = -Max-1
func myAtoi(str string) int {
	var l = len(str)
	if l == 0 {
		return 0
	}
	var begin,end = 0,l-1
	var negative = false
	var hasSign = false
	var nums = make([]uint8,0)
	for begin <= end {
		for begin <= end && (str[begin] == ' ')  {
			begin++
		}
		if begin <= end && str[begin] == '-' {
			negative =  true
			hasSign = true
			begin++
		}
		if !hasSign && begin <= end && str[begin] == '+' {
			hasSign = true
			begin++
		}
		if begin <= end && (str[begin] > '9' || str[begin] < '0') {
			return 0
		}
		for begin <= end && str[begin] == '0' {
			begin++
		}
		for begin <= end && str[begin] <= '9' && str[begin] >= '0' {
			nums = append(nums,str[begin] - '0')
			begin++
		}
		break
	}
	var length = len(nums)
	var res,j = 0,0
	for i:=length-1;i>=0;i--{
		var times = int(math.Pow10(j))
		var r = times * int(nums[i])
		res += r
		j++
		if times > int(Max) || int(Max) < r || res > int(Max){
			if negative {
				return int(Min)
			}
			return int(Max)
		}
	}
	if negative {
		res = -res
	}
	return res
}
