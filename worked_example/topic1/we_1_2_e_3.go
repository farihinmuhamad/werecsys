// "id": "we_1_2_e_3",
// "type": "worked_example",
// "topic": "1. Variables, Data Types, and Operators",
// "difficulty": "medium",
// "context": "Network Latency",
// "title": "Ping Jitter Warning",
// "description": "Check network connection status for ping and jitter. Declare latency metrics, define limits using multi variable declaration, and evaluate status using logical operators."
package main

import "fmt"

func main() {
	// TODO: Declare an integer variable 'pingRate' and assign it a value
	var pingRate int = 118
	// TODO: Declare an integer variable 'jitterRate' and assign it a value
	var jitterRate int = 129
	// TODO: Declare two integer variables 'maxPing' and 'maxJitter' using multi variable declaration and assign them values
	maxPing, maxJitter := 109, 123
	// TODO: Declare a boolean variable 'isLagging' without initial value
	var isLagging bool
	// TODO: Determine 'isLagging' status using logic: True if pingRate greater than maxPing AND jitterRate greater than maxJitter
	isLagging = (pingRate > maxPing) && (jitterRate > maxJitter)
	// TODO: Print the value of isLagging variable
	fmt.Println(isLagging)
}