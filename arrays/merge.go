package arrays

import "sort"

/*
题目:给出一个区间的集合，请合并所有重叠的区间。

解题思路:将所有区间根据start从小到大排序,然后将第一个区间写入返回数组中
然后挨个将后续区间与上一个区间比较

*/
type region struct {
	start,end int
}

type regions []region

func (r regions) Len() int {
	return len(r)
}
func (r regions) Less(i,j int) bool {
	return r[i].start < r[j].start
}
func (r regions) Swap(i,j int) {
	r[i],r[j] = r[j],r[i]
}

func merge(intervals [][]int) [][]int {
	var l = len(intervals)
	if l <= 1 {
		return intervals
	}
	var result = make([][]int,0,l)
	var regions = make(regions,0,l)
	for i := 0 ;i<l;i++{
		regions = append(regions,region{
			start: intervals[i][0],
			end:   intervals[i][1],
		})
	}
	sort.Sort(regions)
	result = append(result,[]int{regions[0].start,regions[0].end})
	for i:=1;i<l;i++ {
		last := result[len(result)-1]
		//当前区间的左边界比上一个区间的右边界大
		if regions[i].start > last[1] {
			result = append(result,[]int{regions[i].start,regions[i].end})
		//当前区间的右边界大于上一个区间的右边界
		}else if regions[i].end > last[1] {
			last[1] = regions[i].end
			result[len(result)-1] = last
		}
	}
	return result
}

/*
给你两个有序整数数组 nums1 和 nums2，请你将 nums2 合并到 nums1 中，使 num1 成为一个有序数组。
双指针:

*/
func mergeArray(nums1 []int, m int, nums2 []int, n int)  {
	var res = make([]int,0,m+n)
	var p,q = 0,0
	for p < m && q < n {
		if nums1[p] < nums2[q] {
			res = append(res,nums1[p])
			p++
		}else{
			res = append(res,nums2[q])
			q++
		}
	}
	if p < m {
		res = append(res,nums1[p:]...)
	}
	if q < n {
		res = append(res,nums2[q:]...)
	}
	copy(nums1,res)
}
