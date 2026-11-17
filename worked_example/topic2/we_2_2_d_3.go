// "id": "we_2_2_d_3",
// "type": "worked_example",
// "topic": "2. Branching (Conditional Statements)",
// "difficulty": "medium",
// "context": "App Menu Selector",
// "title": "Check App Menu Route",
// "description": "Select the correct menu page based on an integer menuOption value. Note: deliberate differences applied — uses switch without expression, and comparison operators inside the cases."
package main

import "fmt"

func main() {
	// TODO: Declare an integer variable 'menuOption' and assign it a value
	var menuOption int = 20
	// TODO: Implement a switch statement without an expression
	switch {
	// TODO: Case statement using comparison operators to check if 'menuOption' is equal to 10
	case menuOption == 10:
		// TODO: Print the Main Menu page message
		fmt.Println("Main Menu")
	// TODO: Case statement to check if 'menuOption' is equal to 20
	case menuOption == 20:
		// TODO: Print the Settings Menu page message
		fmt.Println("Settings Menu")
	// TODO: Case statement to check if 'menuOption' is equal to 30
	case menuOption == 30:
		// TODO: Print the Help Menu page message
		fmt.Println("Help Menu")
	// TODO: Use default for determining Default Case for other options
	default:
		// TODO: Print the Exit App message
		fmt.Println("Exit App")
	}
}
