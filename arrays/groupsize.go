package arrays

import "sort"

/*
给定一副牌，每张牌上都写着一个整数。
此时，你需要选定一个数字 X，使我们可以将整副牌按下述规则分成 1 组或更多组：
每组都有 X 张牌。
组内所有的牌上都写着相同的整数。
仅当你可选的 X >= 2 时返回 true。
*/
func hasGroupsSizeX(deck []int) bool {
	sort.Ints(deck)
	var l = len(deck)
	if l <= 1 {
		return false
	}
	var count = make([]int,0,l/2+1)
	var c,cnt = deck[0],1
	for i:=1;i<l;i++{
		if deck[i] == c {
			cnt++
		}else{
			if cnt == 1 {
				return false
			}
			count = append(count,cnt)
			c = deck[i]
			cnt = 1
		}
	}
	count = append(count,cnt)
	//求最大公约数
	l = len(count)
	if l == 1 {
		return true
	}
	var g = gcd(count[0],count[1])
	for i:=2;i<len(count);i++ {
		g = gcd(g,count[i])
	}
	if g > 1 {
		return true
	}
	return false
}

func gcd(a,b int) int{
	if b == 0 {
		return a
	}
	return gcd(b,a%b)
}
