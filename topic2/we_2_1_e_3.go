// "id": "we_2_1_e_3",
// "type": "worked_example",
// "topic": "2. Branching (Conditional Statements)",
// "difficulty": "easy",
// "context": "Season Classification",
// "title": "Check if a Month is Summer or Winter",
// "description": "Classify a given month as Summer or Winter. Note: deliberate differences applied — uses switch without expression, and comparison/logical operators inside the case statement."
package main

import "fmt"

func main() {
	// TODO: Declare a string variable 'currentMonth' and assign it a value
	var currentMonth string = "July"
	// TODO: Use switch statement without expression
	switch {
	// TODO: Case statement using comparison and logical OR operators
	case currentMonth == "June" || currentMonth == "July":
		// TODO: Print the "Summer" season message
		fmt.Println("Summer")
	// TODO: Use default for determining Default Case for other months
	default:
		// TODO: Print the "Winter" season message
		fmt.Println("Winter")
	}
}
