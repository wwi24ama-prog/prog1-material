package fragen

import "fmt"

type Person struct {
	Name string
	Age  int
	/* ... */
}

func SetAge(p *Person, age int) {
	p.Age = age
}

func ExampleSetAge() {
	p := Person{"Max Mustermann", 0}

	SetAge(&p, 42)

	fmt.Println(p)

	// Output:
}
