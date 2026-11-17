// "id": "we_3_1_b_1",
// "type": "worked_example",
// "topic": "3. Looping",
// "difficulty": "easy",
// "context": "Database Backup",
// "title": "Backup Progress Logging",
// "description": "Simulate database backup progress logging by printing a status message until a float64 limit is reached."
package main

import "fmt"

func main() {
	// TODO: Declare a type inference string variable 'copyBlockMsg' and assign it a value
	copyBlockMsg := "Copying backup block"
	// TODO: Declare a float64 variable 'backupBlockCount' and assign it a value
	var backupBlockCount float64 = 0.0
	// TODO: Declare a type inference float64 variable 'maxBlocks' and assign it a value
	maxBlocks := 46.0
	// TODO: Create a for loop with only a condition to check if 'backupBlockCount' is less than 'maxBlocks'
	for backupBlockCount < maxBlocks {
		// TODO: Print the 'copyBlockMsg' to the console
		fmt.Println(copyBlockMsg)
		// TODO: Increment 'backupBlockCount' by 1 manually inside the loop
		backupBlockCount++
	}
}