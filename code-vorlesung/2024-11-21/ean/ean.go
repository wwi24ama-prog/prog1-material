package ean

// InputOk prüft, ob code 13 Stellen hat
// und nur aus Ziffern besteht.
func InputOk(code string) bool {
	return false
}

// IsDigit erwartet ein Zeichen c und prüft,
// ob dieses eine Ziffer ist.
func IsDigit(c byte) bool {
	return false
}

// ToIntList erwartet einen String "code",
// der nur aus Ziffern besteht.
// Liefert eine Liste dieser Ziffern als []int.
func ToIntList(code string) []int {
	// TODO
	return []int{}
}

// CheckSum erwartet eine Liste von Ziffern und
// berechnet deren EAN-Prüfsumme.
func CheckSum(digits []int) int {
	// TODO
	return 0
}

// EanOk erwartet einen String "code" und liefert true,
// genau dann wenn code eine gültige EAN-Nummer ist.
func EanOk(code string) bool {
	if !InputOk(code) {
		return false
	}

	digits := ToIntList(code)

	checksum := CheckSum(digits)

	return checksum == 0
}
