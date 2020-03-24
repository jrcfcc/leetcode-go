package strchar

/*
字符串有三种编辑操作:插入一个字符、删除一个字符或者替换一个字符。
给定两个字符串，编写一个函数判定它们是否只需要一次(或者零次)编辑。

算法：
首先判断两个字符串长度，相差大于一返回 false
双指针遍历两个字符串，同时记录编辑次数 op_cnt：
若 first[i] == second[j]，不需编辑，i，j 加一
若 first[i] != second[j]，分为三种情况：
first[i] == second[j+1]，那么 j++，op_cnt++
first[i+1] == second[j]，那么 i++，op_cnt++
以上两种都不符合，那么使用替换操作，i++，j++，op_cnt++
注意，一旦 op_cnt > 1，返回 false
遍历结束后，若仍有一方未走到结尾，且相差的长度 + op_cnt 大于 1，则返回 false
*/
func oneEditAway(first string, second string) bool {
	var l1,l2 = len(first),len(second)
	if l1 - l2 > 1 || l2 - l1 > 1 {
		return false
	}
	var i,j,opCount = 0,0,0
	for i < l1 && j < l2 {
		if first[i] == second[j] {
			i++
			j++
			continue
		}
		//删除second[j]
		if j+1<l2 && first[i] == second[j+1] {
			j++
			opCount++
		//删除first[i]
		}else if i+1 < l1 && first[i+1] == second[j] {
			i++
			opCount++
		//替换
		}else{
			i++
			j++
			opCount++
		}
		if opCount > 1 {
			return false
		}
	}
	if l2 - j + opCount > 1 {
		return false
	}
	if l1 - i + opCount > 1 {
		return false
	}
	return true
}
