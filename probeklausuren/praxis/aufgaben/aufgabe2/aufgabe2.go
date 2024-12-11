package aufgabe2

/*
AUFGABENSTELLUNG: Vervollständigen Sie die Funktion ExcludeStringsBetween.
MAX. PUNKTE: 10
*/

// ExcludeStringsBetween erwartet eine Liste und zwei Strings first und last.
// Die Funktion liefert eine Liste mit allen Elementen, die nicht zwischen first und last liegen.

// first und last sollen nicht zum Ergebnis gehören.
// Falls die Liste first oder last nicht enthält, oder falls last vor first vorkommt,
// soll die leere Liste geliefert werden.
func ExcludeStringsBetween(list []string, first, last string) []string {
	var1 := -1
	var2 := -1
	for i := 0; i < len(list); i++ {
		if list[i] == first {
			var1 = i
		}
		if list[i] == last {
			var2 = i
		}

	}
	if var2 <= var1 {
		return []string{}
	}

	return append(list[:var1], list[var2+1:]...)
}
