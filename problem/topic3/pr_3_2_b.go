// "id": "pr_3_2_b",
// "type": "problem",
// "topic": "3. Looping",
// "difficulty": "medium",
// "context": "Data Sync Status",
// "title": "Sync Retry Monitor",
// "description": "Monitor database synchronization status by printing a sync progress state a certain number of times using a while-like for loop and Regular Manifest Typing Variable Declarations. Declare a string constant for sync status, two integer variables for current attempt and limit attempts, and use a for loop to print the sync status until the current attempt reaches the limit.",
// "answer_options": "const syncStatus string = "Synchronizing"; var currentAttempt, limitAttempts int = 0, 42; for currentAttempt < limitAttempts {; currentAttempt++; for currentAttempt <= limitAttempts {; currentAttempt + 1"
package main

import "fmt"

func main() {
	// TODO: Declare a string variable 'syncStatus' and assign it a value
	const syncStatus string = "Synchronizing"
	// TODO: Declare two integer variables 'currentAttempt' and 'limitAttempts' and assign them the values
	var currentAttempt, limitAttempts int = 0, 42
	// TODO: Create a for loop with only a condition to check if 'currentAttempt' is less than 'limitAttempts'
	for currentAttempt < limitAttempts {
		fmt.Println(syncStatus)
		// TODO: Increment 'currentAttempt' by 1 manually inside the loop
		currentAttempt++
	}
}