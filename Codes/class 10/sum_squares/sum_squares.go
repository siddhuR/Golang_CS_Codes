/*

Find the difference between the square of the sum and the sum of the squares of the first N natural numbers.

The square of the sum of the first ten natural numbers is (1 + 2 + ... + 10)² = 55² = 3025.

The sum of the squares of the first ten natural numbers is 1² + 2² + ... + 10² = 385.

Hence the difference between the square of the sum of the first ten natural numbers and the sum of the squares of the first ten natural numbers is 3025 - 385 = 2640.

You are not expected to discover an efficient solution to this yourself from first principles; research is allowed, indeed, encouraged. Finding the best algorithm for the problem is a key skill in software engineering.

*/

package main

import "fmt"

func main() {
	// N is the number of natural numbers
	N := 10

	// Calculate the sum of the squares
	sumOfSquares := 0
	for i := 1; i <= N; i++ {
		sumOfSquares += i * i
	}

	// Calculate the square of the sum
	squareOfSum := 0
	for i := 1; i <= N; i++ {
		squareOfSum += i
	}
	squareOfSum *= squareOfSum

	// Calculate the difference
	difference := squareOfSum - sumOfSquares

	fmt.Printf("The difference between the square of the sum and the sum of the squares of the first %d natural numbers is %d\n", N, difference)
}
