package functional

import "fmt"

// Löscht alle Vokale aus s und liefert das Ergebnis.
func FilterVowels(s string) string {
	if len(s) == 0 {
		return ""
	}

	if !isVowel(s[0]) {
		// Filtere alle Vokale aus dem restlichen String (ab Stelle 1)
		// Liefere das erste Element aus s plus den gefilterten Reststring.
		return s[:0] + FilterVowels(s[1:])
		//return string(s[0]) + FilterVowels(s[1:])
	}
	return FilterVowels(s[1:])
}

func Filter(s string, p func(c byte) bool) string {
	if len(s) == 0 {
		return ""
	}

	if !p(s[0]) {
		return s[:0] + Filter(s[1:], p)
	}
	return Filter(s[1:], p)
}

func ExampleFilterVowels() {
	s1 := "xayezi"

	r1 := FilterVowels(s1)

	fmt.Println(r1)

	// Output:
	// xyz
}
