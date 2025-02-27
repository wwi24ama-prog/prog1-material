package eanm

// InputOk prüft, ob der Code(string) 13 Stellen hat und nur aus Ziffern besteht?
func InputOk(code string) bool {
	for i := 0; i < len(code); i++ { // for _, c := range code
		var c byte = code[i] //string is eine Liste von Bytes
		if !IsDigit(c) {
			return false
		}
	}
	if len(code) == 13 { //return len(code) == 13
		return true
	} else {
		return false
	}
}

// Is Digit erwartet ein Zeichen c (deswegen byte) und prüft ob dieses eine Ziffer ist
func IsDigit(c byte) bool {
	if c >= '0' && c <= '9' {
		return true
	} else {
		return false
	}
}

// ToIntList erwartet einen string ("code"),
// der nur aus Ziffern besteht und Liefert		//result := []int{} <- leere liste erstellen mit länge 0
// eine Liste ([]) dieser Ziffern als []int
func ToIntList(code string) []int { //result := make([]int, len(code)) wenn ich make schreibe kann ich auch vorauswahl bekommen
	list := make([]int, len(code))   // for i := 0; i < len(code); i++ {
	for i := 0; i < len(code); i++ { // 	result [i] = int (code[i] - 'o')
		list[i] = int(code[i] - '0') //}
	}
	return list
}

// CheckSum erwartet eine Liste von Ziffern und berechnet deren EAN-Prüfsumme.
func CheckSum(digits []int) int {
	ergebnis := 0
	for i := 0; i < len(digits); i++ { //for i, n := range digits {}
		n := digits[i]
		if i%2 == 0 {
			ergebnis += n

		} else {
			ergebnis += 3 * n
		}
	}
	return (10 - (ergebnis % 10)) % 10
}

// EanOk erwaretet einen String "code" und liefert true,
// genau dann wenn code eine gültige ean nummer ist
func EanOk(code string) bool {
	if !InputOk(code) { //if InputOk(code) == false
		return false
	}

	digits := ToIntList(code)
	checksum := CheckSum(digits)

	return checksum == 0
}
