package nums

import "strings"

/*中文字符串转整形数字*/

/*
处理一亿以下的数字
*/
func chinesetonums(str string) int {
	var res = 0
	strs := strings.Split(str,"万")
	res = turn(strs[0])
	if len(strs) == 2 {
		res *= 10000
		res += turn(strs[1])
	}
	return res
}

func turn(str string) int {
	var zero = '零'
	var power = make(map[rune]int)
	power['千'] = 1000
	power['百'] = 100
	power['十'] = 10
	var nums = make(map[rune]int)
	nums['一'] = 1
	nums['二'] = 2
	nums['三'] = 3
	nums['四'] = 4
	nums['五'] = 5
	nums['六'] = 6
	nums['七'] = 7
	nums['八'] = 8
	nums['九'] = 9
	var res,num = 0,0
	for _,v := range str {
		if v == zero {
			continue
		}
		if pow,ok:=power[v];ok{
			if num == 0 {
				num = 1
			}
			res = res + num * pow
			num = 0
		}
		if n,ok := nums[v];ok{
			num = n
		}
	}
	res += num
	return res
}
