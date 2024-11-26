package pointer

import "fmt"

type Uhrzeit struct {
	Stunde  int
	Minute  int
	Sekunde int
}

// Tick zählt die Sekunde um 1 hoch.
func Tick(u *Uhrzeit) {
	u.Sekunde++
}

// Tack zählt die Sekunde um 1 hoch.
func (u *Uhrzeit) Tack() {
	u.Sekunde++
}

// Tock zählt auch die Sekunde hoch und liefert das Ergebnis als neues Struct.
func Tock(u Uhrzeit) Uhrzeit {
	u.Sekunde++
	return u
}

func ExampleTick() {
	u := Uhrzeit{0, 0, 0}

	Tick(&u)
	u = Tock(u)
	u.Tack()

	fmt.Println(u)

	// Output:

}
