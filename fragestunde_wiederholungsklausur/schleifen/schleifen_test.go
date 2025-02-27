package schleifen

import "fmt"

func Example() {
	l1 := []int{30, 5, 17, 23, 42, 38}

	// Schleife, die etwas mit der gesamten Liste macht:
	// Immer range.
	for i, v := range l1 {
		fmt.Printf("%v: %v \n", i, v)
	}

	// Schleife, die das erste Element ignoriert:
	// beides möglich
	for i, v := range l1[1:] {
		fmt.Printf("%v: %v \n", i, v)
	}
	for i := 1; i < len(l1); i++ {
		fmt.Printf("%v: %v\n", i, l1[i])
	}

	// Schleife über die geraden Positionen
	// geht nur mit Zähler
	for i := 0; i < len(l1); i += 2 {
		fmt.Println(l1[i])
	}

	// Output:
}
