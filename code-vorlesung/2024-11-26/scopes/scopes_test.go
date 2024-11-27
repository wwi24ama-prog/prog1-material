package scopes

import "fmt"

func Example_scopes() {
	x := 1
	{
		x := 2
		{
			x := 3
			fmt.Println(x)
		}
		fmt.Println(x)
	}
	fmt.Println(x)

	// Output:
}

func Example_scopes_loop() {

	m := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	for i := 0; i < 3; i++ {
		for i := 0; i < 3; i++ {
			fmt.Println(m[i][i])
		}
	}

	// Output:
}

func Foo(x int) {
	fmt.Println(x * 4)
	Bar(x + 2)
}

func Bar(x int) { // x == 44
	Baz(x / 2) // x == 44
	fmt.Println(x)
}

func Baz(x int) { // x == 22
	x += 17 // x == 39
	fmt.Println(x)
}

func ExampleFoo() {
	Foo(42)

	// Output:
	// 168
	// 39
	// 44
}

func Foo1(x int) int {
	if x == 0 {
		return 1
	}
	return 2 * Foo1(x-1)
}

func Example_scopes_rec() {
	fmt.Println(Foo1(0))
	fmt.Println(Foo1(1))
	fmt.Println(Foo1(2))
	fmt.Println(Foo1(3))

	// Output:
	// 1
	// 2
	// 4
	// 8
}
