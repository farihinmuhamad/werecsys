// "id": "pr_2_1_c",
// "type": "problem",
// "topic": "2. Branching (Conditional Statements)",
// "difficulty": "easy",
// "context": "Temperature Check",
// "title": "Check if Temperature is Hot, Normal, or Cold",
// "description": "Write a program that checks if a given temperature is hot, normal, or cold based on predefined thresholds. If the temperature is less than or equal to the cold threshold, print 'Cold'. If the temperature is greater than the cold threshold but less than or equal to the normal threshold, print 'Normal'. If the temperature is greater than the normal threshold, print 'Hot'.",
// "answer_options": "coldThreshold := 20; if temperature <= coldThreshold {; } else if temperature <= normalThreshold {; } else {; { else }; if temperature >= 30 {"
package main

import "fmt"

func main() {
	// TODO: Declare Type Inference integer variable 'coldThreshold' and assign it a value
	coldThreshold := 20
	normalThreshold := 30
	temperature := 19
	// TODO: Create an if statement to check if temperature is Cold
	if temperature <= coldThreshold {
		fmt.Println("Cold")
	// TODO: Create an else if statement to check if temperature is Normal
	} else if temperature <= normalThreshold {
		fmt.Println("Normal")
	// TODO: Create an else statement to handle the case when temperature is Hot
	} else {
		fmt.Println("Hot")
	}
}