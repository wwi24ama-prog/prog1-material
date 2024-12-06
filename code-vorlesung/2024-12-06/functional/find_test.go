package functional

import "fmt"

func FindElement(s string, a byte) int {
	return Find(s, func(c byte) bool {
		return c == a
	})

	// if s == "" {
	// 	return -1
	// }

	// if s[0] == a {
	// 	return 0
	// }

	// return FindElement(s[1:], a) + 1
}

func FirstA(s string) int {
	return Find(s, func(c byte) bool {
		return c == 'a'
	})

	// if s == "" {
	// 	return -1
	// }

	// if s[0] == 'a' {
	// 	return 0
	// }

	// return FirstA(s[1:]) + 1
}

func isVowel(c byte) bool {
	return c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u'
}

func FirstVowel(s string) int {
	return Find(s, isVowel)

	// if s == "" {
	// 	return -1
	// }

	// if s[0] == 'a' || s[0] == 'e' || s[0] == 'i' || s[0] == 'o' || s[0] == 'u' {
	// 	return 0
	// }

	// return FirstVowel(s[1:]) + 1
}

func Find(s string, p func(c byte) bool) int {
	if s == "" {
		return -1
	}

	if p(s[0]) {
		return 0
	}

	return Find(s[1:], p) + 1
}

func ExampleFind() {
	s1 := "xcfuabcde"

	fmt.Println(FirstVowel(s1))
	fmt.Println(FirstA(s1))
	fmt.Println(FindElement(s1, 'x'))

	// Output:
	// 3
	// 4
	// 0
}
