package find_list_meine

import "fmt"

func ExampleFindElementBinary() {
	list := []int{1, 3, 5, 7, 9}

	fmt.Println(FindElementBinary(list, 5))
	fmt.Println(FindElementBinary(list, 9))
	fmt.Println(FindElementBinary(list, 43))

	//Output:
	//2
	//4
	//-1
}
