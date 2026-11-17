// "id": "pr_2_3_b",
// "type": "problem",
// "topic": "2. Branching (Conditional Statements)",
// "difficulty": "hard",
// "context": "Logistics",
// "title": "Shipping Cost",
// "description": "Calculate shipping cost based on weight and destination using nested if logic. Use the following criteria: If the package is international and weighs more than 10 units, apply a high rate. If it's domestic and weighs more than 15 units, also apply a high rate. Otherwise, use a regular rate.",
// "answer_options": "if international {; if weight > 10 {; } else {; if weight > 15 {; if international := true; if weight = 10 {"
package main

import "fmt"

func main() {
	var weight, highRate, regRate int = 15, 25000, 50000
	var international bool = true
	// TODO: Create an if statement to check if 'international' is true
	if international {
		// TODO: Create an if statement to check if 'weight' is greater than 10
		if weight > 10 {
			shippingCost := (highRate * weight) * 2
			fmt.Println("High international shipping cost:", shippingCost)
		} else {
			shippingCost := regRate * weight
			fmt.Println("Regular international shipping cost:", shippingCost)
		}
	//TODO: Create an else statement to handle the case when 'international' is false
	} else {
		// TODO: Create an if statement to check if 'weight' is greater than 15
		if weight > 15 {
			shippingCost := (highRate * weight) * 2
			fmt.Println("High Domestic shipping cost:", shippingCost)
		} else {
			shippingCost := regRate * weight
			fmt.Println("Regular Domestic shipping cost:", shippingCost)
		}
	}
}