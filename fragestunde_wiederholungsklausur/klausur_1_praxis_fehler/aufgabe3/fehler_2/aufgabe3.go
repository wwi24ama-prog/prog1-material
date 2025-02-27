package aufgabe3

// CountX erwartet eine Liste von Zahlen sowie eine Zahl x.
// CountX liefert die Anzahl der Vorkommen von x in der Liste.
func CountX(list []int, x int) int {

	counter := 0

	if len(list) == 0 { // Rekursionsanker: Abbruchbedingung (wenn die Liste = 0 ist)
		return 0
	}

	if list[0] == x { // Prüfen, ob das Slice aus Int am Index 0, die Zahl x enthält
		counter++
		return CountX(list[1:], x)
	}

	if list[0] != x {
		return CountX(list[1:], x)
	}

	return counter
}
