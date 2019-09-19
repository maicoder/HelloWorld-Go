package my_map

import "testing"

func TestInitMap(t *testing.T) {
	m1 := map[int]int{1: 1, 2: 4, 3: 9}
	t.Log(m1[2])
	t.Logf("len m1=%d", len(m1))
	m2 := map[int]int{}
	m2[4] = 16
	t.Logf("len m2=%d", len(m2))
	m3 := make(map[int]int, 10)
	t.Logf("len m3=%d", len(m3))
	var m4 map[int]string
	var m5 = map[int]string{}
	//m4[1] = "2"
	t.Log(m4, m5)
	t.Logf("%p, %p", m4, m5)
	t.Logf("%T, %T", m4, m5)
	if m4 == nil {
		t.Log("m4 == nil")
	}
	if m5 == nil {
		t.Log("m5 == nil")
	}
}

func TestAccessNotExistingKey(t *testing.T) {
	m1 := map[int]int{}
	t.Log(m1[1])
	m1[2] = 0
	t.Log(m1[2])
	m1[3] = 0
	if v, ok := m1[3]; ok {
		t.Logf("Key 3's value is %d", v)
	} else {
		t.Log("Key 3's valus is not existing")
	}
}

func TestTraveMap(t *testing.T) {
	m1 := map[int]int{1: 1, 2: 4, 3: 9}
	for k, v := range m1 {
		t.Log(k, v)
	}
}
