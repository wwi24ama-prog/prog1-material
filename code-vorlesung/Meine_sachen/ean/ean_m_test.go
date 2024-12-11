package eanm

import "fmt"

func ExampleInputOk() {
	fmt.Println(InputOk("1231231231231"))
	fmt.Println(InputOk("1231231231"))
	fmt.Println(InputOk("123123123123A"))
	// Output:
	// true
	// false
	// false
}

// byte mit '
func ExampleIsDigit() {
	fmt.Println(IsDigit('0'))
	fmt.Println(IsDigit('9'))
	fmt.Println(IsDigit('a'))

	//Output:
	//true
	//true
	//false
}

func ExampleToIntList() {
	fmt.Println(ToIntList("1234567891123"))
	//Output:
	//[1 2 3 4 5 6 7 8 9 1 1 2 3]
}

func ExampleCheckSum() {
	fmt.Println(CheckSum([]int{9, 0, 9, 9, 9, 9, 9, 5, 4, 3, 2, 1, 7}))

	// Output:
	// 0
}
