package strchar

import "strings"

func delsubstr(str1,str2 string) string {
	var res = str1
	for strings.Contains(res,str2) {
		res = strings.ReplaceAll(res,str2,"")
	}
	return res
}
