package slice_test

import "testing"

func TestSliceInit(t *testing.T) {
	var s0 []int
	t.Log(s0, len(s0), cap(s0))
	s0 = append(s0, 1)
	t.Log(s0, len(s0), cap(s0))

	s1 := []int{1, 2, 3, 4}
	t.Log(s1, len(s1), cap(s1))

	s1 = append(s1, 5, 6, 7, 8, 9)
	t.Log(s1, len(s1), cap(s1))

	s2 := make([]int, 3, 5)
	t.Log(s2, len(s2), cap(s2))
	s2 = append(s2, 6)
	t.Log(s2, len(s2), cap(s2))
}

func TestSliceGrowing(t *testing.T) {
	s := []int{}
	for i := 0; i < 10; i++ {
		s = append(s, i)
		t.Log(cap(s), len(s))
	}
}

func TestSliceShareMemory(t *testing.T) {
	year := []string{"Jan", "Feb", "May", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"}
	Q2 := year[3:6]
	t.Log(Q2, len(Q2), cap(Q2))
	summer := year[5:8]
	t.Log(summer, len(summer), cap(summer))
	summer[0] = "Unknown"
	t.Log(Q2)
	t.Log(year)
}

func TestSliceCompare(t *testing.T) {
	//a := []int{1, 2, 3, 4}
	//b := []int{1, 2, 3, 4}
	//if a == b {
	//	t.Log("equal")
	//}
}
