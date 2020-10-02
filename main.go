package main

import (
	"sort"
)

func getSkyline(buildings [][]int) [][]int {
	var ret [][]int
	if len(buildings) == 0 {
		return ret
	}

	points := make([][]int, 0)

	for i := 0; i < len(buildings); i++ {
		b := buildings[i]
		p1 := []int{b[0], 1, i}
		p2 := []int{b[1], 2, i}
		points = append(points, p1)
		points = append(points, p2)
	}

	// first sort inner
	sort.SliceStable(points, func(i, j int) bool { return points[i][1] < points[j][1] })
	// next sort first
	sort.SliceStable(points, func(i, j int) bool { return points[i][0] < points[j][0] })

	bit := make(BIT, len(points)+1)

	mp := make(map[int]int)
	for i := 0; i < len(points); i++ {
		if mp[points[i][0]] == 0 {
			mp[points[i][0]] = i
		}
	}

	prevHeight := -1
	for i := 0; i < len(points); i++ {
		pt := points[i]
		var L int
		var R int
		var H int
		if pt[1] == 1 {
			// start of a building
			building := buildings[pt[2]]
			L = building[0]
			R = building[1]
			H = building[2]
			add(bit, mp[R], H)
		} else {
			L = buildings[pt[2]][1]
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
