package strchar

/*
题目:
输入一个字符串，打印出该字符串中字符的所有排列。
你可以以任意顺序返回这个字符串数组，但里面不能有重复元素。

解题思路,回溯算法
*/
var result []string
func permutationUniq(s string) []string {
	var l = len(s)
	result = make([]string,0,2 * l)
	if l <= 1 {
		result = append(result,s)
		return result
	}
	var track = make([]byte,0,l)
	backtrack([]byte(s),track)
	return result
}

func backtrack(b []byte,track []byte) {
	var l = len(b)
	if l == 0 {
		return
	}
	var m = make(map[byte]struct{})
	for i:=0;i<l;i++{
		//去重
		if _,ok:=m[b[i]];ok{
			continue
		}
		m[b[i]] = struct{}{}
		//做选择,将选择项从选择列表中去除
		track = append(track,b[i])
		newB := make([]byte,l-1)
		copy(newB[:i],b[:i])
		if i < l - 1 {
			copy(newB[i:],b[i+1:])
		}
		//回溯
		backtrack(newB,track)
		if len(track) == cap(track) {
			path := make([]byte,len(track))
			copy(path,track)
			result = append(result,string(path))
			return
		}
		//撤销选择
		track = track[:len(track)-1]
	}
}
