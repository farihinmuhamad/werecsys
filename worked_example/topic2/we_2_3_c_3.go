// "id": "we_2_3_c_3",
// "type": "worked_example",
// "topic": "2. Branching (Conditional Statements)",
// "difficulty": "hard",
// "context": "Fintech Account Promotion",
// "title": "Check Account Promotion Status",
// "description": "Apply promotion interest boosts based on account balance, premium membership, and perk availability. Note: deliberate differences applied — reordered branching (non-premium checks first) and inverted internal boolean conditions."
package main

import "fmt"

func main() {
	// TODO: create three variables: isPremium (bool), accountBalance (int), and hasPerk (bool) and assign them appropriate values
	var isPremium bool = true
	var accountBalance int = 140000
	var hasPerk bool = true
	// TODO: Create an if statement to check if isPremium is false and accountBalance is greater than 500000
	if !isPremium && accountBalance > 500000 {
		// TODO: Create another if statement to check if hasPerk is false
		if !hasPerk {
			fmt.Println("Standard Offer Applied")
		} else {
			fmt.Println("Standard Offer + Perk Applied")
		}
		// TODO: Create an else-if statement to check if isPremium is true and accountBalance is greater than 250000
	} else if isPremium && accountBalance > 250000 {
		// TODO: Create another if statement to check if hasPerk is false
		if !hasPerk {
			fmt.Println("Premium Offer Applied")
		} else {
			fmt.Println("Premium Offer + Perk Applied")
		}
		// TODO: Create an else statement to print No Campaign Active if none of the conditions are met
	} else {
		fmt.Println("No Campaign Active")
	}
}
