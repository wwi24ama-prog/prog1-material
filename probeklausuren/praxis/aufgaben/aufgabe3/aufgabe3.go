package aufgabe3

/*
AUFGABENSTELLUNG: Vervollständigen Sie die Funktion CountOdd.
MAX. PUNKTE: 10
ZUSATZBEDINGUNG: Die Funktion muss rekursiv sein.
*/

// CountOdd erwartet eine Liste von Zahlen und liefert die Anzahl der ungeraden Zahlen darin.
func CountOdd(list []int) int {
	if len(list) == 0 {
		return 0
	}
	teiler, speicher := list[0], CountOdd(list[1:])
	if teiler%2 != 0 {
		speicher++
	}
	return speicher
}
