package strchar

/*
判断字符是否唯一

位运算
设置一个int型0作为标记变量,4个字节也就是32个0,
*/
func isUnique(astr string) bool {
   var mark = 0
   for _,v := range astr {
   		//计算该字符的偏移量,然后将1左偏移对应的偏移量
		offset := v - 'a'
		//如果对应偏移位是1,说明这个字符出现过
		if mark & (1 << offset) != 0 {
			return false
		}else{
			//出现过的字符,标记对应的偏移位为1
			mark = mark | (1 << offset)
		}
   }
   return true
}
