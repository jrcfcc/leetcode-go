package strchar


//寻找n个字符串的公共前缀
/*
解题思路1:
假设s1,和s2的最长公共前缀为LCP(s1,s2),
那么n个子字符串的公共前缀为
LCP(s1....sn) = LCP(LCP(LCP(LCP(s1,s2),s3)......sn-1),sn)
*/
func longestCommonPrefix(strs []string) string {
	var n = len(strs)
	if n <= 0 {
		return ""
	}
	//默认第一个字符串为最长的公共前缀
	var pre = strs[0]
	//遍历字符串数组,将其它的字符串与公共前缀比较
	for i:=1;i<n;i++{
		//遍历目前的最长公共前缀
		for j:=0;j< len(pre);j++ {
			//如果最长公共前缀比当前要比较的字符串strs[j]长,截断为当前字符串的长度
			if j >= len(strs[i]) {
				pre = pre[:j]
				continue
			}
			//如果最长公共前缀的字符与要比较的字符串的下标字符不等,将最长公共前缀截断
			if pre[j] != strs[i][j] {
				if j == 0 {
					return ""
				}
				pre = pre[:j]
			}
		}
	}
	return pre
}
