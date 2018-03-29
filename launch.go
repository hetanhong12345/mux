package main

import (
	"fmt"
	"strconv"
)

func main() {
	nums := []int{1, 1, 1, 2, 2, 3}

	fmt.Printf("%#v", nums)
	numDecodings("123")

	fmt.Printf("%t\n", canJump([]int{1, 2, 0, 1}))

}

func numDecodings(s string) int {
	if (len(s) < 1) {
		return 0
	}
	if s[0] == '0' {
		return 0
	}

	dp := make([]int, len(s)+1)
	dp[0] = 0
	dp[1] = 1 //第一位 1,2334

	for j := 2; j <= len(s); j++ {
		dp[j] = 0
		if s[j-1] != '0' {
			//最后一位单独解密
			dp[j] += dp[j-1]

		}
		// 最后两位一起解密
		temp, _ := strconv.Atoi(s[j-2:j])
		if temp >= 10 && temp < 27 {
			if (j == 2) {
				dp[j] += 1
			} else {
				dp[j] += dp[j-2]
			}

		}

	}
	fmt.Printf("%#v\n", dp)
	return dp[len(s)]
}

func canJump(nums []int) bool {

	max := 0
	for index, num := range nums {
		if max < index {
			return false
		}
		if num+index > max {
			max = num + index
		}

	}

	return true

}

func insert(intervals []Interval, newInterval Interval) []Interval {

	var res []Interval
	if len(intervals) == 0 {
		return append(res, newInterval)
	}
	// 插入不在任何区间之内
	if newInterval.End < intervals[0].Start {
		res = append(res, newInterval)
		res = append(res, intervals...)
		return res
	}
	if newInterval.Start > intervals[len(intervals)-1].End {
		return append(intervals, newInterval)
	}
	index := 0
	for index < len(intervals) {
		if intervals[index].End < newInterval.Start && intervals[index+1].Start > newInterval.End {
			res1 := intervals[:index+1]
			res2 := append([]Interval{}, intervals[index+1:]...)
			res = append(res1, newInterval)
			res = append(res, res2...)
			return res
		}
		index++
	}

	tempInterval := newInterval
	for _, item := range intervals {

		// 不在区间内
		if item.Start > tempInterval.End {
			res = append(res, item)
			continue

		}
		// 不在区间内
		if (item.End < tempInterval.Start) {
			res = append(res, item)
			continue

		}
		if len(res) == 0 {
			res = append(res, intervals[0])
		}
		if (tempInterval.Start > item.Start) {
			tempInterval.Start = item.Start

		}
		if tempInterval.End < item.End {
			tempInterval.End = item.End
		}
		last := res[len(res)-1]
		if (last.End < tempInterval.Start) {
			res = append(res, tempInterval)
		} else {
			res[len(res)-1] = tempInterval
		}

	}
	return res

}

func merge1(nums1 []int, m int, nums2 []int, n int) {

	i, j, k := m-1, n-1, m+n-1
	for i >= 0 && j >= 0 {
		if nums1[i] > nums2[j] {
			nums1[k] = nums1[i]
			i--
			k--

		} else {
			nums1[k] = nums2[j]
			j--
			k--
		}
		if (j < 0) {
			return
		}

	}

	for j >= 0 {
		nums1[k] = nums2[j]
		k--
		j--
	}

}

func numTrees(n int) int {
	res := make([]int, n+1)
	res[0], res[1] = 1, 1
	i := 2
	for i <= n {
		for j := 1; j <= i; j++ {
			res[i] = res[j-1] + res[i-j]
		}
		i++

	}
	return res[n]

}
