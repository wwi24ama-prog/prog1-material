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
	result := []string{}

	first_visited := false
	last_visited := false

	for _, el := range list {
		if last_visited && !first_visited {
			return []string{}
		}
		if el == first {
			first_visited = true
		}
		if first_visited == last_visited {
			result = append(result, el)
		}
		if el == last {
			last_visited = true
		}
	}

	if !first_visited || !last_visited {
		return []string{}
	}

	return result
}
