package main

import (
	"fmt"
	"sort"
)

func main() {
	matrix := [][]int{{1, 2}, {3, 4}}
	result := spiralOrder(matrix)
	for _, r := range result {
		fmt.Println(r)
	}
	fmt.Println(generateMatrix(5))
	intervals := []Interval{{1, 6}, {1, 6}, {8, 10}, {15, 18}}
	merge(intervals)
	nums := []int{1, 0, 0, 1, 2, 1, 0}
	sortColors(nums)
	fmt.Printf("%#v", nums)
}
func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return []int{}
	}
	if len(matrix[0]) == 0 {
		return []int{}
	}
	return move(matrix, 0, 0, 0, []int{})

}

/*
r->移动方向 0:→ ，1：↓，2：←，3：↑
i，j-> 当前移动到的位置
result-> 返回值
*/
func move(matrix [][]int, r, i, j int, result []int) ([]int) {
	m := len(matrix)
	n := len(matrix[0])
	result = append(result, matrix[i][j])
	isArrive := 1 << 32
	matrix[i][j] = isArrive;
	// to right
	if r == 0 {
		if (j < n-1 && matrix[i][j+1] != isArrive) {
			return move(matrix, 0, i, j+1, result)
		}
		if (i < m-1 && matrix[i+1][j] != isArrive) {
			return move(matrix, 1, i+1, j, result)
		}
	}
	// to bottom
	if r == 1 {
		if (i < m-1 && matrix[i+1][j] != isArrive) {
			return move(matrix, 1, i+1, j, result)
		}
		if (j > 0 && matrix[i][j-1] != isArrive) {
			return move(matrix, 2, i, j-1, result)
		}
	}
	// to left
	if r == 2 {
		if (j > 0 && matrix[i][j-1] != isArrive) {
			return move(matrix, 2, i, j-1, result)
		}
		if (i > 0 && matrix[i-1][j] != isArrive) {
			return move(matrix, 3, i-1, j, result)
		}
	}
	// to top
	if r == 3 {
		if (i > 0 && matrix[i-1][j] != isArrive) {
			return move(matrix, 3, i-1, j, result)
		}
		if (j < m-1 && matrix[i][j+1] != isArrive) {
			return move(matrix, 0, i, j+1, result)
		}
	}
	return result

}
func generateMatrix(n int) [][]int {

	matrix := [][]int{}
	if n == 1 {
		return [][]int{{1}}
	}
	for index := 0; index < n; index++ {
		ts := make([]int, n)
		matrix = append(matrix, ts)
	}

	i := 0
	j := 0
	count := 0
	var r int // 0:→ ，1：↓，2：←，3：上
	r = 0
	for count < n*n {
		count++
		matrix[i][j] = count
		//fmt.Printf("i=%d ,j=%d,count=%d,matrix[i][j]=%d\n", i, j, count, matrix[i][j])
		if r == 0 {
			if j < n-1 && matrix[i][j+1] == 0 {
				j++
				r = 0
				continue
			}
			if i < n-1 && matrix[i+1][j] == 0 {
				i++
				r = 1
				continue
			}

		}
		if r == 1 {
			if i < n-1 && matrix[i+1][j] == 0 {
				i++
				r = 1
				continue
			}
			if j > 0 && matrix[i][j-1] == 0 {
				j--
				r = 2
				continue
			}

		}
		if r == 2 {
			if j > 0 && matrix[i][j-1] == 0 {
				j--
				r = 2
				continue
			}
			if i > 0 && matrix[i-1][j] == 0 {
				i--
				r = 3
				continue
			}
		}
		if r == 3 {
			if i > 0 && matrix[i-1][j] == 0 {
				i--
				r = 3
				continue
			}
			if j < n-1 && matrix[i][j+1] == 0 {
				j++
				r = 0
				continue
			}
		}
		//fmt.Printf("i=%d ,j=%d,count=%d,matrix[i][j]=%d\n", i, j, count, matrix[i][j])

	}
	return matrix

}

type Interval struct {
	Start int
	End   int
}

func merge(intervals []Interval) []Interval {
	results := []Interval{}
	if len(intervals) == 0 {
		return results
	}
	sort.Slice(intervals[:], func(i, j int) bool {
		return intervals[i].Start < intervals[j].Start
	})
	var temp Interval
	temp = intervals[0]
	results = append(results, temp)
	for _, item := range intervals[1:] {
		if item.Start > temp.End {
			temp = item
			results = append(results, temp)

		}
		if item.End > temp.End {
			(temp).End = (item).End
			results[len(results)-1] = temp

		}

		fmt.Printf("%#v\n", results)

	}

	return results

}

func sortColors(nums []int) {
	i := 0
	j := 0
	for k, _ := range nums {
		temp := nums[k];

		//assigning the current as max
		nums[k] = 2;

		// if original is less than 2 then it should be 1
		if (temp < 2) {
			nums[j] = 1;
			j += 1;
		}

		// if original is less than 1 then it should be 0. The above if statement ensures that 1 - index
		// will always be greater than 0 - index
		if (temp < 1) {
			nums[i] = 0;
			i += 1;
		}
	}

}
