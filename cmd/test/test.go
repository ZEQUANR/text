package main

import (
	"fmt"
	"reflect"
)

func main() {

	const len = 10
	var a = [len + 10]int{1, 2, 4}

	fmt.Println(reflect.TypeOf(a))
}
