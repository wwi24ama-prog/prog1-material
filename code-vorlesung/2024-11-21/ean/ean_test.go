package ean

import "fmt"

func ExampleInputOk() {
	fmt.Println(InputOk("1231231231231"))
	fmt.Println(InputOk("12312312"))
	fmt.Println(InputOk("123123123123A"))

	// Output:
	// true
	// false
	// false
}

func ExampleIsDigit() {
	fmt.Println(IsDigit('0'))
	fmt.Println(IsDigit('9'))
	fmt.Println(IsDigit('a'))

	// Output:
	// true
	// true
	// false
}
