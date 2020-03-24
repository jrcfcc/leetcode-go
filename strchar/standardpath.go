package strchar

import "strings"

/*
题目描述:以 Unix 风格给出一个文件的绝对路径，你需要简化它。或者换句话说，将其转换为规范路径。
在 Unix 风格的文件系统中，一个点（.）表示当前目录本身；此外，两个点 （..） 表示将目录切换到上一级（指向父目录）；两者都可以是复杂相对路径的组成部分。更多信息请参阅：Linux / Unix中的绝对路径 vs 相对路径
请注意，返回的规范路径必须始终以斜杠 / 开头，并且两个目录名之间必须只有一个斜杠 /。最后一个目录名（如果存在）不能以 / 结尾。此外，规范路径必须是表示绝对路径的最短字符串。

解题思路:
1.此题主要考察的是栈,所以定义一个辅助栈;
2.先把字符串以"/"为分隔符分割成数组,此时数组有"路径"、""、"."、".."这四种情况;
3.遍历数组,当s[i].equals("..")并且栈不空时pop,
当!s[i].equals("") && !s[i].equals(".") && !s[i].equals(".."),
即s[i]是路径入栈;
4.栈空,返回"/",栈非空,用StringBuffer做一个连接返回即可;
5完结。
*/
func simplifyPath(path string) string {
	var curr,point,doublePoint = `/`,`.`,`..`
	//字符串模拟栈
	var stack = make([]string,0)
	s := strings.Split(path,curr)
	for _,char := range s {
		if char == doublePoint && len(stack) > 0{
			stack = stack[:len(stack)-1]
		}
		if char != point && char != doublePoint && char != "" {
			stack = append(stack,char)
		}
	}
	if len(stack) == 0 {
		return curr
	}
	var rtn = ""
	for _,str := range stack {
		tmp := "/" + str
		rtn += tmp
	}
	return rtn
}