// interface 参数
package main

import (
	"fmt"
	"strconv"
)

type Human struct {
	name  string
	age   int
	phone string
}

func (h Human) String() string {
	return "(" + h.name + " - " + strconv.Itoa(h.age) + " years - ✆ " + h.phone + ")"
}

func main() {
	Bob := Human{
		name:  "Bob",
		age:   30,
		phone: "111-222-333",
	}
	fmt.Println("This Human is :", Bob)
}
