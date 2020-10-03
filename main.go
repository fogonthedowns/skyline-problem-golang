package main

import (
	"sort"
)

// left and right
// X-axis of building
const LEFTSIDE = 1
const RIGHTSIDE = 2

func getSkyline(buildings [][]int) [][]int {
	var ret [][]int
	if len(buildings) == 0 {
		return ret
	}

	points := make([]Point, 0)

	// [x-axis (value), [1 (left) | 2 (right)], index (building number)]
	for i := 0; i < len(buildings); i++ {
		b := buildings[i]
		left := Point{xAxis: b[0], side: LEFTSIDE, index: i}
		right := Point{xAxis: b[1], side: RIGHTSIDE, index: i}
		points = append(points, left)
		points = append(points, right)
	}

	// first sort inner
	sort.SliceStable(points, func(i, j int) bool { return points[i].side < points[j].side })
	// next sort first
	sort.SliceStable(points, func(i, j int) bool { return points[i].xAxis < points[j].xAxis })

	bit := make(BIT, len(points)+1)

	// maps all points on the x-axis
	mp := make(map[int]int)
	for i := 0; i < len(points); i++ {
		if mp[points[i].xAxis] == 0 {
			mp[points[i].xAxis] = i
		}
	}

	prevHeight := -1
	for i := 0; i < len(points); i++ {
		pt := points[i]
		var L int
		var R int
		var H int

		if pt.side == LEFTSIDE {
			// start of a building
			building := buildings[pt.index]
			L = building[0]
			R = building[1]
			H = building[2]
			add(bit, mp[R], H)
		} else {
			L = buildings[pt.index][1]
		}

		cur := find(bit, mp[L]+1)
		if cur != prevHeight {
			if len(ret) > 0 {
				if ret[len(ret)-1][0] == L {
					ret[len(ret)-1][1] = max(ret[len(ret)-1][1], cur)
				} else {
					ret = append(ret, []int{L, cur})
				}
			} else {
				ret = append(ret, []int{L, cur})
			}
			prevHeight = cur
		}
	}

	return ret

}

type BIT []int
type Point struct {
	xAxis int
	side  int
	index int
}

// builds the BIT, i increases with the opp i -= i & (-i)
// uses max(cur, height)
func add(b BIT, i int, h int) {
	for i > 0 {
		b[i] = max(b[i], h)
		i -= i & (-i)
	}
}

func find(b BIT, L int) int {
	ret := 0
	for L <= len(b) {
		ret = max(ret, b[L])
		L += L & (-L)
	}
	return ret
}

func max(i, j int) int {
	if i > j {
		return i
	} else {
		return j
	}
}
