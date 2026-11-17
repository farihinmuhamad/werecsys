// "id": "pr_2_1_e",
// "type": "problem",
// "topic": "2. Branching (Conditional Statements)",
// "difficulty": "easy",
// "context": "Day Classification",
// "title": "Determine Weekend or Weekday",
// "description": "Write a program that determines if a given day is a weekend or a weekday. Declare a string variable 'day' and assign it a value, then use a switch statement to check its value, print 'Weekend' for Saturday and Sunday, and print 'Weekday' for all other days.",
// "answer_options": "var day string = "Monday"; switch day {; case "Sunday", "Saturday":; default:; switch { Sunday; case default"
package main

import "fmt"

func main() {
	// TODO: Declare a string variable 'day' and assign it a value
	var day string = "Monday"
	// TODO: Use switch statement to check the 'day' variable
	switch day {
	// TODO: Case statement to check if it's a weekend (Sunday or Saturday)
	case "Sunday", "Saturday":
		fmt.Println("Weekend")
	// TODO: Use default for determining Default Case for other days
	default:
		fmt.Println("Weekday")
	}
}