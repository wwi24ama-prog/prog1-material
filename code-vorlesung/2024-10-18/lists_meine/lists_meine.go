package lists_meine

import "fmt"

func ListDemo() {

	//Definiere eine Liste von int.
	list1 := []int{1, 3, 5, 7, 9} //[]=Liste int=von ganzen Zahlen {}=Zahlen in Liste mit ,

	//Gibt das erste Element auf der Konsole aus
	fmt.Println(list1[0])
	fmt.Println(list1[4])

	//Gib ein Element an ungültiger Stelle aus
	//fmt.Println(list1[5])

	//Gib alle Element der Liste aus

	fmt.Println()

	for i := 0; i < len(list1); i++ { //len=Laenge
		fmt.Println(list1[i])
	}

}
