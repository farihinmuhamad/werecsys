// "id": "we_1_3_b_3",
// "type": "worked_example",
// "topic": "1. Variables, Data Types, and Operators",
// "difficulty": "hard",
// "context": "Network Rate Limiting",
// "title": "Traffic Load Verification",
// "description": "Evaluate if traffic level is within network request limits using multi-variable declaration of integers, constants, blank identifier, and logical operator checks."
package main

import "fmt"

func main() {
	// TODO: Declare an integer constant 'MaxRequestLimit' and assign it a value
	const MaxRequestLimit int = 22000
	// TODO: Declare type inference Multi Variable 'ratePerSecond' and 'activeDuration' (int) and assign it a value
	ratePerSecond, activeDuration := 131, 149
	// TODO: Declare a string variable 'clientIP' and assign it a value
	clientIP := "10.0.0.1"
	// TODO: Declare an Underscore '_' variable and assign it the value of the 'clientIP' variable
	_ = clientIP
	// TODO: Declare a boolean variable 'hasSecurityThreat' and assign it a value
	hasSecurityThreat := false
	// TODO: Declare a Type Inference variable 'isSafeTraffic' (bool) to check if 'ratePerSecond' multiplied by 'activeDuration' is less than 'MaxRequestLimit' AND 'hasSecurityThreat' is false
	isSafeTraffic := (ratePerSecond*activeDuration < MaxRequestLimit) && !hasSecurityThreat
	// TODO: Print the value of the 'isSafeTraffic' variable
	fmt.Println(isSafeTraffic)
}