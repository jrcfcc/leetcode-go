package offer

/*
替换空格
请实现一个函数，把字符串 s 中的每个空格替换成"%20"。
预先申请一个byte数组,大小为s的三倍,然后从后向前扫描
*/
func replaceSpace(s string) string {
	var l = len(s)
	var res = make([]byte,3*l)
	var j = 3*l-1
	for i:=l-1;i>=0;i--{
		if s[i] == ' '{
			res[j] = '0'
			res[j-1] = '2'
			res[j-2] = '%'
			j-=2
		}else{
			res[j] = s[i]
		}
		j--
	}
	return string(res[j+1:])
}
