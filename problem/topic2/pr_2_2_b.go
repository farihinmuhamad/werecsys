// "id": "pr_2_2_b",
// "type": "problem",
// "topic": "2. Branching (Conditional Statements)",
// "difficulty": "medium",
// "context": "Discount Eligibility",
// "title": "Purchase Discount Evaluation",
// "description": "Evaluate the total purchase amount and print the corresponding discount using if-else statements. Use the following discount scale: 10% for 250,000 and above, 5% for 100,001 to 2499,999, and no discount for 100,000 and below.",
// "answer_options": "const minSpending int = 250000; totalSpending := 200000; if totalSpending >= minSpending {; } else if totalSpending > 100000 {; } else {; var totalSpending float64 = 175000"
package main

import "fmt"

func main() {
	// TODO: Declare a constant 'minSpending' of type int and assign it a value
	const minSpending int = 250000
	// TODO: Declare a Type Inference integer variable 'totalSpending' and assign it a value
	totalSpending := 200000
	// TODO: Create an if statement to check if totalSpending is greater than or equal to minSpending
	if totalSpending >= minSpending {
		fmt.Println("Discount 10%")
	// TODO: Else if statement to check if totalSpending is greater than 100000
	} else if totalSpending > 100000 {
		fmt.Println("Discount 5%")
	} else {
		fmt.Println("No Discount")
	}
}