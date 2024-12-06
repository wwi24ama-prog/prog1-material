package aufgabe6

/*
AUFGABENSTELLUNG: Vervollständigen Sie die Funktion SymmetricDifference.
MAX. PUNKTE: 10
*/

// SymmetricDifference erwartet zwei int-Listen l1 und l2.
// Die Funktion liefert eine int-Liste mit den Elementen,
// die in einer, aber nicht in beiden Listen vorhanden sind.
//
// Die Elemente kommen dabei in der gleichen Reihenfolge vor, wie in den
// Ursprungslisten, mehrfach vorkommende Elemente werden entsprechend wiederholt.
// Die Elemente aus l1 kommen vor denen aus l2 in der Ergebnisliste vor.
func SymmetricDifference(l1, l2 []int) []int {
	ergebnis := []int{}

	for i := 0; i < len(l1); i++ {
		zergebnis := l1[i]
		found := false
		for j := 0; j < len(l2); j++ {
			if zergebnis == l2[j] {
				found = true
				break
			}

		}
		if !found {
			ergebnis = append(ergebnis, zergebnis)
		}

		for i := 0; i < len(l2); i++ {
			zergebnis := l2[i]
			found := false
			for j := 0; j < len(l1); j++ {
				if zergebnis == l1[j] {
					found = true
					break
				}
			}
			if !found {
				ergebnis = append(ergebnis, zergebnis)
			}
		}

	}
	return ergebnis
}
