// "id": "pr_3_3_c",
// "type": "problem",
// "topic": "3. Looping",
// "difficulty": "hard",
// "context": "Log Filtering",
// "title": "Critical Event Locator",
// "description": "Search a series of logs for a critical status. Print a search message, loop through iterations, skip normal logs with continue, and stop when a critical status is found using break. Declare variables for search mode, log count, and critical index; use a for loop to iterate through logs, print search mode and log number, use continue to skip normal logs, and break when the critical log is reached.",
// "answer_options": "for i := 1; i <= logCount; i++ { | if i < criticalIndex { | if i == criticalIndex { | break | for i = 1; i < logCount; i+1 { | if i != criticalIndex {"
package main

import "fmt"

func main() {
	var searchMode string = "Inspect"
	var logCount, criticalIndex int = 37, 35
	// TODO: Create a regular for loop to iterate 'i' from 1 to 'logCount' while incrementing 'i' by 1 at each iteration
	for i := 1; i <= logCount; i++ {
		// TODO: Create an if statement to check if 'i' is less than 'criticalIndex'
		if i < criticalIndex {
			continue
		}
		fmt.Println(searchMode, i)
		// TODO: Create an if statement to check if 'i' is equal to 'criticalIndex'
		if i == criticalIndex {
			// TODO: Execute 'break' statement to exit the loop when 'i' is equal to 'criticalIndex'
			break
		}
	}
}