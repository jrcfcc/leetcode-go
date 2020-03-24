package strchar

import "strconv"

/*
给定一个只包含数字的字符串，复原它并返回所有可能的 IP 地址格式。

回溯算法
backtrack(选择列表,路径)
for 选择 in 选择列表
	做选择
	将选择加入路径
	backtrack(选择列表,路径)
    撤销选择

选择三个位置将`.`插入到字符串中
*/
var l int
var str string
var res []string
func restoreIpAddresses(s string) []string {
	l = len(s)
	str = s
	res = make([]string,0,30)
	var track = make([]byte,0,l+3)
	ipbacktrack(-1,3,track)
	return res
}

func ipbacktrack(pre_pos,dots int,track []byte) {
	var max_pos = l - 1
	if pre_pos + 4 < max_pos {
		max_pos = pre_pos + 4
	}
	//做选择,可选范围在上一个点的位置pre+1->pre+4范围内
	for curr_pos := pre_pos + 1;curr_pos < max_pos;curr_pos++ {
		segment  := str[pre_pos+1:curr_pos+1]
		if isValid(segment) {
			//本次选择的字符串长度
			var ls = len(segment)
			track = append(track,segment...)
			track = append(track,'.')
			//这是最后一个点,判断最后的部分是否符合ip的要求
			if dots - 1 == 0{
				segment = str[curr_pos+1:]
				if isValid(segment) {
					var ls = len(segment)
					track = append(track,segment...)
					//加入到结果集
					res = append(res,string(track))
					//撤销选择
					track = track[:len(track)-ls]
				}
			}else{
				ipbacktrack(curr_pos,dots-1,track)
			}
			//撤销选择
			track = track[:len(track)-ls-1]
		}
	}
}

//判断是否可以插入点
func isValid(segment string) bool {
	var l = len(segment)
	if l > 3{
		return false
	}
	v,e := strconv.Atoi(segment)
	if e != nil {
		return false
	}
	if segment[0] == '0' && l == 1 {
		return true
	}
	if segment[0] != '0' && v <= 255{
		return true
	}
	return false
}



/*
给你一个有效的 IPv4 地址 address，返回这个 IP 地址的无效化版本。
所谓无效化 IP 地址，其实就是用 "[.]" 代替了每个 "."。
*/
func defangIPaddr(address string) string {
	var l = len(address)
	var result = make([]byte,l+6)
	var j = l + 5
	for i := l-1;i>=0;i--{
		if address[i] == '.' {
			result[j] = ']'
			result[j-1] = '.'
			result[j-2] = '['
			j-=2
		}else{
			result[j] = address[i]
		}
		j--
	}
	return string(result)
}
