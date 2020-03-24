package arrays

//归并排序
func mergesort(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}
	//递归切割成相邻的两部分
	partA := mergesort(nums[:len(nums)/2])
	partB := mergesort(nums[len(nums)/2:])

	//临时数组,用来存放相邻的两部分排好序的部分
	temp := make([]int, len(partA) + len(partB))

	aPointer := 0
	bPointer := 0
	i := 0

	//从小到大移动数据到临时数组
	for aPointer < len(partA) && bPointer < len(partB) {
		if partA[aPointer] < partB[bPointer] {
			temp[i] = partA[aPointer]
			aPointer++
		} else {
			temp[i] = partB[bPointer]
			bPointer++
		}
		i++
	}
	//a还有数据未移动到临时数组
	for aPointer < len(partA) {
		temp[i] = partA[aPointer]
		aPointer++
		i++
	}
	//b还有数据未移动到临时数组
	for bPointer < len(partB) {
		temp[i] = partB[bPointer]
		bPointer++
		i++
	}
	return temp
}

