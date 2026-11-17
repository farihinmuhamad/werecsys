// "id": "we_3_1_a_2",
// "type": "worked_example",
// "topic": "3. Looping",
// "difficulty": "easy",
// "context": "Progressive Accumulator",
// "title": "Geometric Power Accumulator",
// "description": "Accumulate power-of-two squares up to a threshold using a conditional while-style loop."
package main

import "fmt"

func main() {
	// TODO: Declare a type inference integer variable 'maxCapLimit' and assign it a value
	maxCapLimit := 50
	// TODO: Declare a type inference integer variable 'currentNum' and assign it a value
	currentNum := 1
	// TODO: Declare a type inference integer variable 'sumSquares' and assign it a value
	sumSquares := 0
	// TODO: Create a while-style loop that continues as long as 'currentNum' is less than or equal to 'maxCapLimit'
	for currentNum <= maxCapLimit {
		// TODO: Change the value of 'sumSquares' by adding the square of 'currentNum' to it
		sumSquares += currentNum * currentNum
		// TODO: Change the value of 'currentNum' by multiplying it by 2
		currentNum *= 2
	}
	// TODO: Print the value of 'sumSquares' to the console
	fmt.Println(sumSquares)
}