package fragen

import "fmt"

func Example_if() {
	if Foo1() {

	}

	// if 23 || 42 {

	// }

}

func Foo1() bool {
	return true
}

func ExampleSetVariable() {
	x := 0
	SetVariable(&x)
	fmt.Println(x)

	// Output:
}

func SetVariable(x *int) {
	*x = 77
	fmt.Println(*x)
}

// Nicht klausurrelevant, nur Frage beantwortet.
func Example_map() {
	m := map[string]int{}

	m["Max Mustermann"] = 42

	fmt.Println(m)
	fmt.Println(m["Max Mustermann"])

	// Output:
}

func Example_defs() {
	n := 42.0

	fmt.Printf("%T", n)

	// Output:
}
