package leetcode

import "testing"

type Data11 struct {
	heigth []int
	expect int
}

func Test11(t *testing.T) {
	ds11 := []Data11{
		{
			heigth: []int{1, 8, 6, 2, 5, 4, 8, 3, 7},
			expect: 49,
		},
		{
			heigth: []int{1, 1},
			expect: 1,
		},
		{
			heigth: []int{4, 3, 2, 1, 4},
			expect: 16,
		},
		{
			heigth: []int{1, 2, 1},
			expect: 2,
		},
	}
	for i, v := range ds11 {
		t.Log(maxArea(v.heigth))
		if maxArea(v.heigth) != v.expect {
			t.Logf("error index: %d \n", i)
		}
	}

}

func TestSort(t *testing.T) {
	T := [][]int{
		{
			1, 8, 6, 2, 5, 4, 8, 3, 7,
		},
		{
			1, 1,
		},
		{
			4, 3, 2, 1, 4,
		},
		{
			1, 2, 1,
		},
	}
	for _, v := range T {
		s := Sort(v, false)
		t.Log(s.nums)
		t.Log(s.index)
	}
}
