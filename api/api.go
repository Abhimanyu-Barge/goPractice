package api

import (
	"fmt"
)

func StartWebServer(Port int) (int, error) {
	fmt.Println("Server is started..")
	// var i int
	// for i := 0; i < 5; i++ {
	// 	println(i)
	// 	// if i == 3 {
	// 	// 	break
	// 	// }
	// }
	var i int
	for {
		if i == 0 {
			break
		}
		println(i)
		i++
	}

	slice := []int{1, 2, 34, 4}
	for index, data := range slice {
		println(index, data)
	}

	mapdemo := map[string]int{"one": 1, "two": 2}
	for k, v := range mapdemo {
		println(k, v)
	}

	// panic

	return Port, nil
}
