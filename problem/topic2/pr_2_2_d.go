// "id": "pr_2_2_d",
// "type": "problem",
// "topic": "2. Branching (Conditional Statements)",
// "difficulty": "medium",
// "context": "Role-Based Access Control",
// "title": "User Role Permissions",
// "description": "Use a switch statement to print user role permissions based on the value of the 'role' integer variable. The program should print the corresponding access level based on the role value: 1 for 'Administrator Access', 2 for 'General User Access', 3 for 'View Only Access', and 'No Access' for any other value."
// "answer_options": "switch role {; case 1:; case 3:; default:; switch {; default role:"
package main

import "fmt"

func main() {
	var role int = 5
	// TODO: Implement a switch statement to handle different 'role' values and print the corresponding access level.
	switch role {
	// TODO: Case statement to check if 'role' is equal to 1
	case 1:
		fmt.Println("Administrator Access")
	case 2:
		fmt.Println("General User Access")
	// TODO: Case statement to check if 'role' is equal to 3
	case 3:
		fmt.Println("View Only Access")
	//TODO: Use default for determining Default Case for other roles
	default:
		fmt.Println("No Access")
	}
}