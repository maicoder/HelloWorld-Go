// method 重写
package main

import "fmt"

type Human struct {
	name  string
	age   int
	phone string
}

type Student struct {
	Human
	school string
}

type Employee struct {
	Human
	company string
}

// Human 定义的 method
func (h *Human) SayHi() {
	fmt.Printf("Hi, I am %s you can call me on %s\n", h.name, h.phone)
}

// Employee 重写的method
func (e *Employee) SayHi() {
	fmt.Printf("Hi, I am %s, I work at %s, you can call me on %s.\n", e.name, e.company, e.phone)
}

func main() {
	mark := Student{
		Human:  Human{"Mark", 25, "111-222-333"},
		school: "MIT",
	}
	sam := Employee{
		Human:   Human{"Sam", 45, "555-666-777"},
		company: "Golang Inc",
	}

	mark.SayHi()
	sam.SayHi()

}
