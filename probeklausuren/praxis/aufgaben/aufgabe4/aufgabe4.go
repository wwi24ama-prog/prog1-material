package aufgabe4

/*
AUFGABENSTELLUNG: Vervollständigen Sie die Funktion MaxElements.
MAX. PUNKTE: 10
ZUSATZBEDINGUNG: Die Funktion muss rekursiv sein!
*/

// MaxElements erwartet zwei int-Listen und liefert eine Liste, die an jeder Position
// jeweils das größere der beiden Elemente enthält.
// Falls eine Position nur in einer Liste vorkommt, gilt dieses Element als das größere.
func MaxElements(l1, l2 []int) []int {
	if len(l1) == 0 {
		return l2
	}
	if len(l2) == 0 {
		return l1
	}
	great := l1[0]

	if l2[0] > great {
		great = l2[0]
	}

	return append([]int{great}, MaxElements(l1[1:], l2[1:])...)
}
