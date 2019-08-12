package pipefilter

import "testing"

func TestSumElems(t *testing.T)  {
	sf := NewSumFilter()
	ret, err := sf.Process([]int{1,2,3})
	if err != nil {
		t.Fatal(err)
	}
	expected := 5
	if ret != expected {
		t.Fatalf("the expected is %d, but actual is %d", expected, ret)
	}
}

func TestWrongInputForSumFilter(t *testing.T)  {
	sf := NewSumFilter()
	_, err := sf.Process([]float32{1.1, 2.2, 3.3 })

	if err == nil {
		t.Fatal("An error is xpected")
	}
}