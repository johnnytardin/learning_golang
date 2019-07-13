package main

import (
	"fmt"
	"reflect"
)

func main() {
	a := [5]int{}
	fmt.Println(reflect.TypeOf(a))
	a[0] = 200
	fmt.Println(a)

	b := make([]int, 10)
	fmt.Println(reflect.TypeOf(b))
	b[1] = 100
	b = append(b, 9)
	fmt.Println(b)

}
