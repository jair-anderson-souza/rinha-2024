package main

import "fmt"

func main() {
	fmt.Print("Test Module")

	text := "text"
	fmt.Println(text)

	var texto = "text"
	fmt.Print(texto)

	var (
		numb    int    = 1
		example string = "asdasd"
	)

	fmt.Println(numb)
	fmt.Println(example)

	variavel5, variavel6 := "variavel 5", "Variavel 6"

	fmt.Println(variavel5)
	fmt.Println(variavel6)

	const PI = "3.14"
	fmt.Println(PI)

}

func Example() {
	main()
}
