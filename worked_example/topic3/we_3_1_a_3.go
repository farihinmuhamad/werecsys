// "id": "we_3_1_a_3",
// "type": "worked_example",
// "topic": "3. Looping",
// "difficulty": "easy",
// "context": "Number Aggregation",
// "title": "Reverse Countdown Sum",
// "description": "Calculate the sum of first N numbers in reverse order using a decreasing loop with a decrement operator. Note: deliberate differences applied — reverse loop with decrement and short variable declaration."
package main

import "fmt"

func main() {
	// TODO: Declare an integer variable 'maxSteps' and assign it a value
	var maxSteps int = 19
	// TODO: Declare an integer variable 'sequenceTotal' and assign it a value
	sequenceTotal := 0
	// TODO: Create a regular for loop to iterate from 'maxSteps' down to 1 while decrementing 'i' by 1 at each iteration
	for i := maxSteps; i >= 1; i-- {
		// TODO: Add the current index 'i' to 'sequenceTotal'
		sequenceTotal += i
	}
	// TODO: Print the value of 'sequenceTotal' to the console
	fmt.Println(sequenceTotal)
}