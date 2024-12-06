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
	result := []int{}

	c := 2

	for n > 1 {
		if n%c == 0 {
			result = append(result, c)
			n = n / c
		} else {
			c++
		}
	}

	return result
}

func PrimeFactorsRecursive(n int) []int {
	return PrimeFactorsRecursiveHelper(n, 2)
}

func PrimeFactorsRecursiveHelper(n, c int) []int {
	if n <= 1 {
		return []int{}
	}

	result := []int{}

	if n%c == 0 {
		result = append(result, c)
		n = n / c
	} else {
		c++
	}

	return append(result, PrimeFactorsRecursiveHelper(n, c)...)

}

func ExamplePrimeFactors() {
	fmt.Println(PrimeFactors(1))
	fmt.Println(PrimeFactors(60))
	fmt.Println(PrimeFactors(42))
	fmt.Println(PrimeFactors(23))
	fmt.Println(PrimeFactorsRecursive(1))
	fmt.Println(PrimeFactorsRecursive(60))
	fmt.Println(PrimeFactorsRecursive(42))
	fmt.Println(PrimeFactorsRecursive(23))

	// Output:
	// []
	// [2 2 3 5]
	// [2 3 7]
	// [23]
	// []
	// [2 2 3 5]
	// [2 3 7]
	// [23]
}
