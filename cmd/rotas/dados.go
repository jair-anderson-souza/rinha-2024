package main

import "fmt"

func testNumbers() int8 {

	var (
		n1 int8 = 111
		n2 int16
		n3 int32
		n4 int64
		n5 int
		n6 uint = 0 //unsigned
	)

	fmt.Println(n1, n2, n3, n4, n5, n6)

	return n1
}
