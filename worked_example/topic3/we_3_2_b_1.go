// "id": "we_3_2_b_1",
// "type": "worked_example",
// "topic": "3. Looping",
// "difficulty": "medium",
// "context": "Queue Processing",
// "title": "Queue Job Logger",
// "description": "Log a job processing status message a set number of times using a while-like loop, a string constant, and a multi variable declaration with float64 counter."
package main

import "fmt"

func main() {
	// TODO: Declare a string constant 'batchLabel' and assign it a value
	const batchLabel string = "Processing job"
	// TODO: Declare two variables 'processedCount' and 'batchLimit' using multi variable declaration
	var processedCount, batchLimit int = 0, 44
	// TODO: Create a for loop with only a condition to check if 'processedCount' is less than 'batchLimit'
	for batchLimit > processedCount {
		// TODO: Print the value of 'batchLabel'
		fmt.Println(batchLabel)
		// TODO: Increment 'processedCount' by 1
		processedCount++
	}
}