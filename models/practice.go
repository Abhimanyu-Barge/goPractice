package models

import "fmt"

const (
	pi    = 3.14159265358979323846
	first = iota
	second
)
const (
	third = iota
)

type Student struct {
	name string
	age  int
}

var (
	User   []*Student
	NextId = 1
)

func PracticeDemo() {
	fmt.Println("hello ...")
	// decleared and assign
	var i int
	i = 12
	fmt.Println(i)
	// decleared assign
	var f float32 = 3.12
	fmt.Println(f)
	// short hand declearatiion
	firstName := "abhimanyu "
	fmt.Println(firstName)
	// complex data type
	c := complex(3, 2)
	fmt.Println(c)
	r, im := real(c), imag(c)
	fmt.Println(r)
	fmt.Println(im)

	//pointer data type
	var name1 *string = new(string)
	fmt.Println(*name1)
	*name1 = "pointer"
	fmt.Println(*name1)

	name := "pointer1"
	ptr := &name
	fmt.Println(ptr, *ptr)
	name = "pointer2"
	fmt.Println(ptr, *ptr)

	// constant
	fmt.Println(pi + 1)
	fmt.Println(pi + 1.4)

	// iota
	fmt.Println(first)
	fmt.Println(second)
	fmt.Println(third)
	// collection type
	// 1)array
	var arr [3]int
	arr[0] = 1
	arr[1] = 2
	arr[2] = 3
	fmt.Println("arr", arr)
	// short hand declrtion
	arr1 := [3]int{1, 2, 3}
	fmt.Println("arr1", arr1)

	// 2) slice
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println(slice)
	fmt.Println("slice declration", slice)
	// copy slice
	slice1 := arr[:]
	slice1[2] = 5
	slice1 = append(slice1, 5)
	slice2 := arr[1:]
	fmt.Println(arr, slice1, slice2)
	// 3) map
	m := map[string]int{}
	m["foo"] = 1
	m["foo1"] = 2
	fmt.Println("map", m)

	// 4) struct

	// init struct
	var user Student
	user.name = "abhi"
	user.age = 12
	fmt.Println("struct", user)
	// short hand declaration
	user2 := Student{name: "abhi", age: 2}
	fmt.Println("userr2", user2)
}
