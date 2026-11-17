// "id": "we_2_3_c_1",
// "type": "worked_example",
// "topic": "2. Branching (Conditional Statements)",
// "difficulty": "hard",
// "context": "E-commerce Reward",
// "title": "Evaluate Customer Purchase Bonuses",
// "description": "Apply customer rewards based on purchase value, registration status, and promo code availability using float64 types."
package main

import "fmt"

func main() {
	// TODO: Declare and initialize isRegistered with type inference
	isRegistered := false
	// TODO: Declare and initialize purchaseValue with float64 type inference
	purchaseValue := 575000.5
	// TODO: Declare and initialize promoAvailable with type inference
	promoAvailable := false
	// TODO: Create an if statement to check if isRegistered is true and purchaseValue is greater than 250000.0
	if isRegistered && purchaseValue > 250000.0 {
		// TODO: Create another if statement to check if promoAvailable is true
		if promoAvailable {
			fmt.Println("Registered Bonus + Promo Added")
		} else {
			fmt.Println("Registered Bonus Added")
		}
		// TODO: Create an else-if statement to check if isRegistered is false and purchaseValue is greater than 500000.0
	} else if !isRegistered && purchaseValue > 500000.0 {
		// TODO: Create another if statement to check if promoAvailable is true
		if promoAvailable {
			fmt.Println("Standard Bonus + Promo Added")
		} else {
			fmt.Println("Standard Bonus Added")
		}
		// TODO: Create an else statement to print No Promotion Applied if none of the conditions are met
	} else {
		fmt.Println("No Promotion Applied")
	}
}
