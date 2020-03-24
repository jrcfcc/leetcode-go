package strchar

/*
给定两个字符串 s1 和 s2，请编写一个程序，确定其中一个字符串的字符重新排列后，
能否变成另一个字符串。

解题思路,统计两个字符串中每个字符出现的频率是否一致
一致则可以
利用辅助数组,统计26个字母的出现频率从a-z分别对应0-25的索引

高级解法,异或运算
异或运算的特点就是两个相等的数异或的结果是0,把两个字符串挨个异或计算一遍,最终结果为0的话
就说明这两个字符串中字符出现的频率是一致的
*/
func CheckPermutation(s1 string, s2 string) bool {
	var l1,l2 = len(s1),len(s2)
	if l1 != l2 {
		return false
	}
	var arr1,arr2 = make([]uint8,26),make([]uint8,26)
	for i:=0;i<l1;i++{
		arr1[s1[i]-'a']++
		arr2[s2[i]-'a']++
	}
	for i:=0;i<=25;i++{
		if arr1[i] != arr2[i] {
			return false
		}
	}
	return true
}

func CheckPermutation1(s1 string, s2 string) bool {
	var l1,l2 = len(s1),len(s2)
	if l1 != l2 {
		return false
	}
	var base = uint8(0)
	for i:=0;i<l1;i++{
		base = base^s1[i]^s2[i]
	}
	if base == 0 {
		return true
	}
	return false
}
