package strchar

/*
题目:
给定两个以字符串形式表示的非负整数 num1 和 num2，返回 num1 和 num2 的乘积，
它们的乘积也表示为字符串形式。

解题思路:
竖式相乘
*/
func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	var result = ""
	for i := len(num1) -1 ;i>=0;i--{
		var carry,res = byte(0),""
		//补0
		for j:=0;j<len(num1)-1-i;j++{
			res += "0"
		}
		var a = num1[i] - '0'
		for j:=len(num2) -1 ;j>=0;j--{
			var b = num2[j] - '0'
			summary := a * b + carry
			res = byte2Str(summary%10) + res
			carry = summary/10
		}
		if carry != byte(0) {
			res = byte2Str(carry) + res
		}
		//两个数组相加
		result = addStrings(result,res)
	}
	return  result
}
