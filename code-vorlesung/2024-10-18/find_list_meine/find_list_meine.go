package find_list_meine

//FindElement erwatet eine Liste von Zahlen und eine weiter Zahl n.
//Die Funktion liefert die Position von n in der Liste.
//Falls n nicht in der Liste vorkommt, liefert die Funktion -1
//
// Die funktion führt eine lineare Suche in der Liste durch
// D.h. sie durchsucht die Liste von Anfang bis Ende.

func FindElementLinear(list []int, n int) int {
	for i := 0; i < len(list); i++ { //laufen durch die Liste. müssen wir bei listen immer machen
		if list[i] == n {
			return i
		}

	}
	return -1
}

//FindElementBinary erwatet eine Liste von Zahlen und eine weiter Zahl n.
//Die Funktion liefert die Position von n in der Liste.
//Falls n nicht in der Liste vorkommt, liefert die Funktion -1
//
// Die funktion führt eine binäre Suche in der Liste durch
// D.h. sie startet in der Mitte und macht dann im linken oder rechten Teil weiter
//Voraussetzung: Die Liste muss sortiert sein!

func FindElementBinary(list []int, n int) int {

	if len(list) == 0 {
		return -1
	}

	m := len(list) / 2
	if list[m] == n { //Test ob es zufällig stimmt wenn ja direkt ende
		return m
	}

	if n < list[m] {
		return FindElementBinary(list[:m], n) //[0:m] Slice ein Ausschnitt der Liste 0-m. Geht von 0-m-1

	}

	result := FindElementBinary(list[m+1:], n) //Er betrachtet nur den Ausschnitt, weswegen es nicht das richtige ergebniss geben würde deswegen +m+1
	if result == -1 {
		return -1
	}
	return result + m + 1
}
