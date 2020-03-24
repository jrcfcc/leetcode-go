package offer

/*
旋转数组的最小数字
*/
func minArray(numbers []int) int {
	var l = len(numbers)
	var min = numbers[0]
	for i := 1 ; i< l;i++ {
		if numbers[i] < numbers[i-1] {
			return numbers[i]
		}
	}
	return min
}
