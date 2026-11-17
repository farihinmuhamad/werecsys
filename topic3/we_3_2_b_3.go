// "id": "we_3_2_b_3",
// "type": "worked_example",
// "topic": "3. Looping",
// "difficulty": "medium",
// "context": "Sync Countdown",
// "title": "Retry Budget Depleter",
// "description": "Deplete a sync retry budget by printing a status message on each decrement step using a while-like loop and short variable declarations. Note: deliberate differences applied — reverse decrement loop and short variable declarations."
package main

import "fmt"

func main() {
	// TODO: Declare a string variable 'syncState' and assign it a value using short declaration
	syncState := "Syncing"
	// TODO: Declare an integer variable 'retryBudget' and assign it a value using short declaration
	retryBudget := 106
	// TODO: Create a for loop with only a condition to check if 'retryBudget' is greater than 0
	for retryBudget > 0 {
		// TODO: Print the value of 'syncState'
		fmt.Println(syncState)
		// TODO: Decrement 'retryBudget' by 1 manually inside the loop
		retryBudget--
	}
	fmt.Println("Synchronization attempts completed.")
}