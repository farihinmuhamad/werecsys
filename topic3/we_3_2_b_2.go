// "id": "we_3_2_b_2",
// "type": "worked_example",
// "topic": "3. Looping",
// "difficulty": "medium",
// "context": "Data Sync Audit",
// "title": "Sync Attempt Logger",
// "description": "Log database synchronization attempt status for each round using a regular 3-clause for loop and a string constant."
package main

import "fmt"

func main() {
	// TODO: Declare a string constant 'syncRound' and assign it a value
	const syncRound string = "Sync attempt"
	// TODO: Declare an integer variable 'syncAttempts' and assign it a value
	var syncAttempts int = 31
	// TODO: Create a regular for loop to iterate 'x' from 1 to 'syncAttempts' while incrementing 'x' by 1
	for x := 1; x <= syncAttempts; x++ {
		// TODO: Print the value of 'syncRound' and the current value of 'x'
		fmt.Println(syncRound, x)
	}
}