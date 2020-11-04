package main

import "fmt"

func main(){
	fmt.Println(insert([][]int{[]int{1,2},[]int{3,5}},[]int{2,5}))
}


func insert(intervals [][]int, newInterval []int) [][]int {
	if len(intervals) == 0{
		return [][]int{newInterval}
	}
	start := newInterval[0]
	end := newInterval[1]
	for _,v := range intervals{
		if v[0] <= start && v[1] >= start{
			start = v[0]
		}

		if v[0] <= end && v[1]>= end {
			end = v[1]
			break
		}

		if v[0] >end {
			break
		}
	}

	var res [][]int
	insert := false
	for i,v:= range intervals{

		if start <= v[0] && end >= v[1] && !insert{
			res = append(res,[]int{start,end})
			insert = true
			continue
		}
		if start < v[0] && end < v[0] {
			if !insert {
				res = append(res, []int{start, end})
				insert = true
			}
			res = append(res,intervals[i:len(intervals)]...)
			break
		}

		if start > v[1] && end >v[1] {
			res = append(res,v)
		}


	}
	if !insert{
		res = append(res, []int{start, end})
	}

	return res
}