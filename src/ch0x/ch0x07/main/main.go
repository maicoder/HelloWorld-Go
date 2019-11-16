package main

import (
	"fmt"
	"strconv"
)

type Element interface {
}

type List []Element

type Person struct {
	name string
	age  int
}

func (p Person) String() string {
	return "(name: " + p.name + "- age: "+ strconv.Itoa(p.age) + "year)"
}

func main() {
	list := make(List, 3)
	list[0] = 1
	list[1] = "hello"
	list[2] = Person{
		name: "Dennis",
		age:  70,
	}

	for index, element := range list {
		switch value := element.(type) {
		case int:
			fmt.Printf("list[%d] is an int and its value is %d\n", index, value)
		case string:
			fmt.Printf("list[%d] is an string and its value is %d\n", index, value)
		}
		
	}

}
