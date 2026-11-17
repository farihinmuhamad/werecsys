// "id": "we_2_1_a_2",
// "type": "worked_example",
// "topic": "2. Branching (Conditional Statements)",
// "difficulty": "easy",
// "context": "Aerospace System",
// "title": "Overpressure Warning Check",
// "description": "Check if boiler pressure has exceeded maximum safety limits. Perform short-statement if initialization to check threshold criteria."
package main

import "fmt"

func main() {
	// TODO: Declare a constant int 'maxPsi' and assign it a value
	const maxPsi int = 182
	// TODO: Perform short statement if initialization and check if 'psiValue' is greater than or equal to 'maxPsi'
	if psiValue := 194; psiValue >= maxPsi {
		// TODO: Print the overpressure alert message
		fmt.Println("Overpressure Alert")
	}
}
