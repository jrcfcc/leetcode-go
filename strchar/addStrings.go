package strchar

/*
题目:给定两个字符串形式的非负整数 num1 和num2 ，计算它们的和。

解题思路:
算法流程： 设定 i，j 两指针分别指向 num1，num2 尾部，模拟人工加法；

计算进位： 计算 carry = tmp // 10，代表当前位相加是否产生进位；
添加当前位： 计算 tmp = n1 + n2 + carry，并将当前位 tmp % 10 添加至 res 头部；
索引溢出处理： 当指针 i或j 走过数字首部后，给 n1，n2 赋值为 0，
相当于给 num1，num2 中长度较短的数字前面填 0，以便后续计算。
当遍历完 num1，num2 后跳出循环，并根据 carry 值决定是否在头部添加进位 11，
最终返回 res 即可。

*/
func addStrings(num1 string, num2 string) string {
	var len1,len2 = len(num1),len(num2)
	var carry,res,i,j = byte(0),"",len1-1,len2-1
	for i >= 0 || j >= 0 {
		a,b := byte(0),byte(0)
		if i >= 0 {
			a = num1[i] - '0'
		}
		if j >= 0 {
			b = num2[j] - '0'
		}
		summary :=  a + b + carry
		res = byte2Str(summary % 10) + res
		carry = summary / 10
		i--
		j--
	}
	if carry != byte(0) {
		res = "1" + res
	}
	return res
}

func byte2Str(a byte) string {
	return string(a + '0')
}
