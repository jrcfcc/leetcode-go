package strchar

import "strings"

/*
给定一个字符串，逐个翻转字符串中的每个单词。
输入字符串可以在前面或者后面包含多余的空格，但是反转后的字符不能包括。
如果两个单词间有多余的空格，将反转后单词间的空格减少到只含一个。

解题思路:
按空格 分割字符串
然后对每个字符串去除空格,然后反转数组
*/
func reverseWords(s string) string {
	a := strings.Split(s," ")
	var result = ""
	l := len(a)
	for i:= l-1;i>=0;i-- {
		b := strings.TrimSpace(a[i])
		if b != "" {
			if result == "" {
				result = b
			}else{
				result = result + " " + b
			}
		}
	}
	return result
}
