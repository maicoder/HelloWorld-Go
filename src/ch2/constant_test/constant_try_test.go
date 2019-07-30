package constant_test

import "testing"

const (
	Monday = iota + 1
	Tuesday
	Wednesday
)

const(
	Readable = 1 << iota
	Writable
	Executable
)

func TestConstantTry(t *testing.T) {
	t.Log(Monday, Tuesday, Wednesday)
}

func TestConstant1(t *testing.T)  {
	a := 1
	t.Log(a&Readable == Readable, a&Writable == Writable, a&Executable == Executable)
}