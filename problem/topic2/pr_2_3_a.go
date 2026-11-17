// "id": "pr_2_3_a",
// "type": "problem",
// "topic": "2. Branching (Conditional Statements)",
// "difficulty": "hard",
// "context": "Railway Transit Control",
// "title": "Train Speed Status Check",
// "description": "Determine train speed status (High, Express, Normal, Slow) based on current speed value using if-else statements. If the speed is greater than or equal to 285, print 'Bullet Train'. If the speed is greater than or equal to 195 but less than 285, print 'Express Train'. If the speed is greater than or equal to 115 but less than 195, print 'Normal Train'. If the speed is less than 115, print 'Slow Train'.",
// "answer_options": "if currentSpeed >= 285 {; } else if currentSpeed >= 195 {; } else if currentSpeed >= 115 {; } else {; else if currentSpeed >= 195; else if currentSpeed >= 115"
package main

import "fmt"

func main() {
	currentSpeed := 225
	// TODO: Create an if statement to check if 'currentSpeed' is greater than or equal to 285
	if currentSpeed >= 285 {
		fmt.Println("Bullet Train")
	// TODO: Create an else-if statement to check if 'currentSpeed' is greater than or equal to 195
	} else if currentSpeed >= 195 {
		fmt.Println("Express Train")
	// TODO: Create another else-if statement to check if 'currentSpeed' is greater than or equal to 115
	} else if currentSpeed >= 115 {
		fmt.Println("Normal Train")
	// TODO: Create an else statement to print "Slow Train" if none of the conditions are met
	} else {
		fmt.Println("Slow Train")
	}
}