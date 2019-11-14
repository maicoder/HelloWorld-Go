package main

import "fmt"

/*
func add1(a int) int {
	a = a+1
	return a
}

func add2(a *int) int {
	*a = *a + 1
	return *a
}
*/

type Human struct {
	name string
	age int
	weight int
}

type Skills []string

type Student struct {
	Human  // 匿名字段，那么就默认Student包含了Human的所有字段
	Skills  // 匿名字段，自定义的类型string slice
	int  // 内置类型作为匿名字段
	speciality string
}

func main() {
	/*
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

	 */

	/*
	mark := Student{Human{"Mark", 25, 120}, "Computer Science"}

	fmt.Println("His name is ", mark.name)
	fmt.Println("His age is ", mark.age)
	fmt.Println("His weight is ", mark.weight)
	fmt.Println("His speciality is ", mark.speciality)

	mark.speciality = "AI"
	fmt.Println("Mark changed his speciality")
	fmt.Println("His speciality is ", mark.speciality)

	fmt.Println("Mark become old")
	mark.age = 46
	fmt.Println("his age is ", mark.age)

	fmt.Println("Mark is not an athlet anymore")
	mark.weight += 60
	fmt.Println("His weight is ", mark.weight)

	mark.Human = Human{
		name:   "Marcus",
		age:    55,
		weight: 220,
	}
	mark.Human.age -= 1
	 */

	jane := Student{
		Human:      Human{"Jane", 35, 100},
		Skills:     nil,
		int:        0,
		speciality: "Biology",
	}

	fmt.Println("Her name is", jane.name)
	fmt.Println("Her age is ", jane.age)
	fmt.Println("Her weight is ", jane.weight)
	fmt.Println("Her speciality is ", jane.speciality)

	jane.Skills = []string{"anatomy"}
	fmt.Println("Her skill is ", jane.Skills)
	fmt.Println("she acquired two new ones")
	jane.Skills = append(jane.Skills, "physics", "golang")
	fmt.Println("Her skills now are ", jane.Skills)

	jane.int = 3
	fmt.Println("Her preferred number is ", jane.int)

}
