// "id": "pr_2_2_c",
// "type": "problem",
// "topic": "2. Branching (Conditional Statements)",
// "difficulty": "medium",
// "context": "Network Bandwidth Allocation",
// "title": "Traffic Priority Classification",
// "description": "Classify network traffic priority based on delay times and packet loss using IF-Else IF-Else conditionals with logical operators. Use the following criteria: 'Drop Connection' if delayMs is greater than 245 OR packetLoss is greater than 23, 'Throttle Priority' if delayMs is greater than 148 AND packetLoss is greater than 9, and 'Normal Priority' for all other cases.",
// "answer_options": "var delayMs, packetLoss int = 165, 11; if delayMs > MaxDelayLimit || packetLoss > 23 {; } else if delayMs > 148 && packetLoss > 9 {; fmt.Println("Normal Priority"); if delayMs > 148 && packetLoss > 9 {; else if delayMs > MaxDelayLimit || packetLoss > 23 {"
package main

import "fmt"

func main() {
	const MaxDelayLimit int = 245
	// TODO: Declare two integer variables 'delayMs' and 'packetLoss' using Multi Variable Declaration and assign them values
	var delayMs, packetLoss int = 165, 11
	// TODO: Create an if statement to check if 'delayMs' is greater than 'MaxDelayLimit' OR 'packetLoss' is greater than 23
	if delayMs > MaxDelayLimit || packetLoss > 23 {
		fmt.Println("Drop Connection")
	// TODO: Create an else-if statement to check if 'delayMs' is greater than 148 AND 'packetLoss' is greater than 9
	} else if delayMs > 148 && packetLoss > 9 {
		fmt.Println("Throttle Priority")
	} else {
		// TODO: Print "Normal Priority" if none of the above conditions are met
		fmt.Println("Normal Priority")
	}
}