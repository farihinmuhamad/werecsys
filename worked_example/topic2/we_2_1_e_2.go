// "id": "we_2_1_e_2",
// "type": "worked_example",
// "topic": "2. Branching (Conditional Statements)",
// "difficulty": "easy",
// "context": "Fruit Classification",
// "title": "Check if a Fruit is Tropical or Temperate",
// "description": "Identify if a given fruit belongs to the tropical or temperate category using a switch statement and string comparison. Declare a string variable, assign it a value, and evaluate it."
package main

import "fmt"

func main() {
	// TODO: Declare a string variable 'selectedFruit' and assign it a value
	var selectedFruit string = "Peach"
	// TODO: Use switch statement to check the 'selectedFruit' variable
	switch selectedFruit {
	// TODO: Case statement to check if it's a tropical fruit (Durian or Papaya)
	case "Durian", "Papaya", "Mango", "Banana":
		// TODO: Print the "Tropical" category message
		fmt.Println("Tropical")
	// TODO: Use default for determining Default Case for other fruits
	default:
		// TODO: Print the "Temperate" category message
		fmt.Println("Temperate")
	}
}
