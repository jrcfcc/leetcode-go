package backtrack

/*
题目:给定一个数组,返回它的全排列

解法:回溯算法
解决一个回溯问题，实际上就是一个决策树的遍历过程。你只需要思考 3 个问题：
1、路径：也就是已经做出的选择。
2、选择列表：也就是你当前可以做的选择。
3、结束条件：也就是到达决策树底层，无法再做选择的条件。
伪码:
result = []
def backtrack(路径, 选择列表):
    if 满足结束条件:
        result.add(路径)
        return

    for 选择 in 选择列表:
        做选择(排除不合法的选择)
        backtrack(路径, 选择列表)
        撤销选择
*/
var result [][]int
func permute(nums []int) [][]int {
	var l = len(nums)
	result = make([][]int,0,2 * l)
	if l <= 1 {
		result = append(result,nums)
		return result
	}
	var track = make([]int,0,l)
	backtrack(nums,track)
	return result
}

//返回不重复的排列
func permuteUnique(nums []int) [][]int {
	var l = len(nums)
	result = make([][]int,0,2 * l)
	if l <= 1 {
		result = append(result,nums)
		return result
	}
	var track = make([]int,0,l)
	backtrack(nums,track)
	return result
}

//nums表示所有可能的选择 track表示当前路径
func backtrack(nums []int,track []int) {
	var length = len(nums)
	if length == 0 {
		return
	}
	var m = make(map[int]struct{})
	for k,v := range nums {
		//去重
		if _,ok:=m[v];ok{
			continue
		}
		m[v] = struct{}{}
		//做选择,也就是把选择加入路径
		//此处应该把track中不存在的元素加入到track中
		//比较简单的操作是做完选择后,把对应的选择从nums中去掉
		track = append(track,v)
		newNum := make([]int,length-1)
		copy(newNum[:k],nums[:k])
		if k < length - 1 {
			copy(newNum[k:],nums[k+1:])
		}
		//回溯
		backtrack(newNum,track)
		//路径中已包含所有选择,结束回溯
		if len(track) == cap(track) {
			path := make([]int,len(track))
			copy(path,track)
			result = append(result,path)
			return
		}
		//撤销选择
		track = track[:len(track)-1]
	}
}


