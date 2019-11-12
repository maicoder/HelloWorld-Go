package main

import "fmt"

func add1(a int) int {
	a = a+1
	return a
}

func add2(a *int) int {
	*a = *a + 1
	return *a
}

func main() {
	x := 3
	y := 13

	fmt.Println("x = ", x)
	fmt.Println("y = ", y)

	x1 := add1(x)
	y1 := add2(&y)

	fmt.Println("x = ", x)
	fmt.Println("x+1=", x1)

	fmt.Println("y = ", y)
	fmt.Println("y + 1", y1)
}
