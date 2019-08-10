package testing

import (
	"fmt"
	"testing"
)

func TestSquare(t *testing.T) {
	inputs := [...]int{1, 2, 3}
	expected := [...]int{1, 4, 9}
	for i := 0; i < len(inputs); i++ {
		ret := square(inputs[i])
		if ret != expected[i] {
			t.Errorf("input is %d, the expected is %d, the actual %d", inputs[i], expected[i], ret)
		}
	}
}

// Fail, Error : 该测试失败，该测试继续， 其他测试继续执行
// FailNow， Fatal： 该测试失败，该测试中止，其他测试继续执行

func TestErrorInCode(t *testing.T) {
	fmt.Println("Start.")
	t.Error("Error")
	fmt.Println("End")
}

func TestFailInCode(t *testing.T) {
	fmt.Println("Start.")
	t.Fatal("Error")
	fmt.Println("End")
}
