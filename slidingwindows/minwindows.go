package slidingwindows

/*
给你一个字符串 S、一个字符串 T，请在字符串 S 里面找出：包含 T 所有字母的最小子串。

示例：

输入: S = "ADOBECODEBANC", T = "ABC"
输出: "BANC"
说明：

如果 S 中不存这样的子串，则返回空字符串 ""。
如果 S 中存在这样的子串，我们保证它是唯一的答案。

滑动窗口算法思路:
滑动窗口算法的思路是这样：
1、我们在字符串 S 中使用双指针中的左右指针技巧，初始化 left = right = 0，
把索引闭区间 [left, right] 称为一个「窗口」。
2、我们先不断地增加 right 指针扩大窗口 [left, right]，
直到窗口中的字符串符合要求（包含了 T 中的所有字符）。
3、此时，我们停止增加 right，转而不断增加 left 指针缩小窗口 [left, right]，
直到窗口中的字符串不再符合要求（不包含 T 中的所有字符了）。
同时，每次增加 left，我们都要更新一轮结果。
4、重复第 2 和第 3 步，直到 right 到达字符串 S 的尽头。
*/
func minWindow(s string, t string) string {
	var res = ""
	var l1,l2 = len(s),len(t)
	if l1 == 0 || l2 > l1 {
		return res
	}
	var max = int(^uint(0)>>1)
	var needs = make(map[rune]int)
	var needsSize = 0 //统计t中不同字符的个数
	//记录t中每个字符出现的次数
	for _,c := range t {
		if count,ok := needs[c];ok{
			count++
			needs[c] = count
		}else{
			needs[c] = 1
			needsSize++
		}
	}
	var window = make(map[rune]int)
	var left,right,match = 0,0,0
	for right < l1 {
		var c = s[right]
		var count,ok = window[rune(c)]
		if ok{
			count++
		}else{
			count = 1
		}
		window[rune(c)] = count
		if count == needs[rune(c)] {
			match++
		}
		right++
		//窗口中已包含t中所有的字符,缩小左边界,寻找最短的满足条件的子串
		for match == needsSize {
			//记录当前满足条件的子字符的长度
			now := right - left
			if now < max {
				res = s[left:right]
				max = now
			}
			c = s[left]
			_,ok = needs[rune(c)]
			//左边界是t的子字符
			if ok {
				window[rune(c)]--
				//滑动窗口中某个子字符数量小于t中的数量
				if window[rune(c)] < needs[rune(c)] {
					match--
				}
			}
			left++
		}
	}
	return res
}
