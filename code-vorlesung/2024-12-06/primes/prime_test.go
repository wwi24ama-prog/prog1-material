package primes

import "fmt"

func IsPrime(n int) bool {
	return n > 1 && !DivisibleByAny(n, n/2)
}

func DivisibleByAny(n, c int) bool {
	if c <= 1 {
		return false
	}

	if n%c == 0 {
		return true
	}
	return DivisibleByAny(n, c-1)
}

func ExampleIsPrime() {
	fmt.Println(IsPrime(1))
	fmt.Println(IsPrime(2))
	fmt.Println(IsPrime(3))
	fmt.Println(IsPrime(4))
	fmt.Println(IsPrime(5))
	fmt.Println(IsPrime(6))

	// Output:
	// false
	// true
	// true
	// false
	// true
	// false
}

func PrimeFactors(n int) []int {
	// TODO
	return []int{}
}

func ExamplePrimeFactors() {
	fmt.Println(PrimeFactors(60))

	// Output:
	// [2 2 3 5]
}
