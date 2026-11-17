// "id": "we_2_1_d_2",
// "type": "worked_example",
// "topic": "2. Branching (Conditional Statements)",
// "difficulty": "easy",
// "context": "Network Packet",
// "title": "Check if Packet Size is Aligned",
// "description": "Check if a network packet size is byte-aligned or misaligned using modulus and a switch statement. Use a short-statement switch initialization to declare and evaluate the packet size."
package main

import "fmt"

func main() {
	// TODO: Use switch statement with short-statement initialization for packetSize
	switch packetSize := 419; packetSize % 2 {
	// TODO: Case statement to check for aligned packet size (remainder 0)
	case 0:
		// TODO: Print "Aligned" for case 0
		fmt.Println("Aligned")
	// TODO: Default case to handle misaligned packet size
	default:
		// TODO: Print "Misaligned" for default case
		fmt.Println("Misaligned")
	}
}
