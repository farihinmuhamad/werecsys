// "id": "pr_3_1_b",
// "type": "problem",
// "topic": "3. Looping",
// "difficulty": "easy",
// "context": "Packet Transmission",
// "title": "Retry Alert Logging",
// "description": "Simulate packet transmission retry logging by printing an alert message a certain number of times using a while-like for loop and Regular Manifest Typing Variable Declaration. Declare a string variable for the alert message and an integer variable for the retry count and a maximum retry limit, then use a loop to print the alert message until the retry count reaches the specified maximum.",
// "answer_options": "var alertMessage string = "Resending packet"; var retryCount int = 0; for retryCount < maxRetries {; retryCount++; if retryCount < maxRetries {; retryCount--"
package main

import "fmt"

func main() {
	// TODO: Declare a string variable 'alertMessage' and assign it a value
	var alertMessage string = "Resending packet"
	// TODO: Declare an integer variable 'retryCount' and assign it a value
	var retryCount int = 0
	var maxRetries int = 46
	// TODO: Create a for loop with only a condition to check if 'retryCount' is less than 'maxRetries'
	for retryCount < maxRetries {
		fmt.Println(alertMessage)
		// TODO: Increment 'retryCount' by 1 manually inside the loop
		retryCount++
	}
}