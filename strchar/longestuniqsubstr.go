package strchar

/*
寻找不重复的最长子字符串的长度
解题思路:滑动窗口
假设子字符串s[i,j)(左闭右开)是没有重复字符的,那么当窗口右滑时,我们只需要判断s[j]是否存在于
子串s[i,j)中,如果不存在,说明s[i,j]是不重复的子串,窗口继续右滑,i不变,j+1,如果s[j]存在于
s[i,j)中那么说明s[i,j)就是其中一个不重复的子串,记录下长度,
把i的下标移动到刚刚重复的那个字符的下一个索引,j不变，窗口继续右滑,继续这个过程,得到所有不重复子串的长度
取最大值即可
窗口以一个map维护
*/
func lengthOfLongestSubstring(s string) int {
	var n = len(s)
	if n < 1 {
		return 0
	}
	var slideMap = make(map[byte]int,n)
	var maxLen = 1
	var left = 0 //本次遍历子串的起始下标
	for i:= 0;i<n;i++ {
		//出现重复字符了,并且重复字符大于起始字符
		//为什么要加上slideMap[s[i]] > left的限定条件
		//因为当出现重复字符,起始下标跳到重复字符的下一个索引时,map中存储的字符并没有清理掉
		//所以子串是否重复要加上重复的字符在起始字符的右边,也就是重复字符要在子串内
		if index,ok := slideMap[s[i]];ok && slideMap[s[i]] >= left {
			//取最长的无重复子串的最大值
			if i - left > maxLen {
				maxLen = i - left
			}
			left = index + 1
		}
		//右滑窗口
		slideMap[s[i]] = i
	}
	//滑到了最右边时,将最后一个未重复的子串长度与最大值比较
	if n - left > maxLen {
		maxLen = n - left
	}
	return maxLen
}
