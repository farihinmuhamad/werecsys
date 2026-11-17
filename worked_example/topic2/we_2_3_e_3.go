// "id": "we_2_3_e_3",
// "type": "worked_example",
// "topic": "2. Branching (Conditional Statements)",
// "difficulty": "hard",
// "context": "Portal Access Control",
// "title": "Verify Portal Access Permissions",
// "description": "Evaluate portal permissions based on user group, validation status, and failed attempts using a switch statement. Note: deliberate differences applied — short declarations and reordered branches."
package main

import "fmt"

func main() {
	// TODO: Declare three type inference variables for userGroup, portalRegistered, and failedAttempts
	userGroup := "admin"
	portalRegistered := true
	failedAttempts := 1
	// TODO: Create a switch statement without an expression
	switch {
	// TODO: Case statement to check if the user has more than 3 failed attempts
	case failedAttempts > 3:
		fmt.Println("Portal Access Denied")
	// TODO: Case statement to check if the user is a viewer and registered
	case portalRegistered && userGroup == "viewer":
		fmt.Println("View Permission Allowed")
	// TODO: Case statement to check if the user is an editor and registered
	case portalRegistered && userGroup == "editor":
		fmt.Println("Edit Permission Allowed")
	// TODO: Case statement to check if the user is an admin and registered
	case portalRegistered && userGroup == "admin":
		fmt.Println("Admin Permission Allowed")
	// TODO: Case statement to check if the user has a known group but is not registered
	case userGroup == "admin" || userGroup == "editor" || userGroup == "viewer":
		fmt.Println("Registration Verification Required")
	// TODO: Default case to handle unrecognized user groups or conditions
	default:
		fmt.Println("Unrecognized Group Level")
	}
}