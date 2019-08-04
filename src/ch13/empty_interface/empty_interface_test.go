package empty_interface

import (
	"fmt"
	"testing"
)

func DoSomething(p interface{}) {
	if i, ok := p.(int); ok {
		fmt.Println("Integer", i)
		return
	}
	if s, ok := p.(string); ok {
		fmt.Println("string", s)
		return
	}
	fmt.Println("Unknown Type")
}

func TestEmptyInterfaceAssertion(t *testing.T) {
	DoSomething(10)
	DoSomething("10")
}
