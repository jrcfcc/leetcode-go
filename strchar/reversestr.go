package strchar

/*
给定一个字符串和一个整数 k，你需要对从字符串开头算起的每个 2k 个字符的前k个字符进行反转。
如果剩余少于 k 个字符，则将剩余的所有全部反转。如果有小于 2k 但大于或等于 k 个字符，
则反转前 k 个字符，并将剩余的字符保持原样
*/

func reverseStr(s string, k int) string {
	var l = len(s)
	if l < 2*k {
		return s
	}
	str := reverse(s[:2*k])
	if l >= 4*k {
		return str
	}
	if l < 3*k {
		remain := reverse(s[2*k:])
		str = str + remain
		return str
	}
	if l >= 3*k {
		remain := reverse(s[2*k:3*k])
		str = str + remain
		return str
	}
	return str
}

func reverse(s string) string {
	res := []byte(s)
	var l = len(res)
	var left,right = 0,l-1
	for left < right {
		res[left],res[right] = res[right],res[left]
		left++
		right--
	}
	return string(res)
}