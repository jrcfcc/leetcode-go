package strchar

import "strconv"

/*
字符串压缩。利用字符重复出现的次数，编写一种方法，实现基本的字符串压缩功能。
比如，字符串aabcccccaaa会变为a2b1c5a3。若“压缩”后的字符串没有变短，则返回原先的字符串。
你可以假设字符串中只包含大小写英文字母（a至z）。

*/
func compressString(S string) string {
	var l = len(S)
	if l <= 1 {
		return S
	}
	var res = make([]byte,0)
	var c,cnt = S[0],1
	for i:=1;i<l;i++ {
		if S[i] == c {
			cnt++
		}else{
			res = append(res,c)
			res = append(res,strconv.Itoa(cnt)...)
			cnt = 1
			c = S[i]
		}
	}
	res = append(res,c)
	res = append(res,strconv.Itoa(cnt)...)
	if len(res) >= l {
		return S
	}
	return string(res)
}
