package learning

import (
	"errors"
	"fmt"
)

func Numbers() {
	var (
		n1 int8 = 111
		n2 int16
		n3 int32
		n4 int64
		n5 int
		n6 uint = 0 //unsigned
	)

	fmt.Println(n1, n2, n3, n4, n5, n6)

	var n7 rune = 234
	fmt.Println(n7)

	var n8 byte = 123
	fmt.Println(n8)

	n9 := 2.3
	var n10 float32 = 6.7
	fmt.Println(n9, n10)

	var texto = "Test"
	fmt.Println(texto)

	//Tabela ASC
	char := "B"
	fmt.Println(char)

	var boleano bool = true
	var err error = errors.New("erro interno")

	fmt.Println(boleano, err)
}
