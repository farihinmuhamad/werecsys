// "id": "pr_2_3_e",
// "type": "problem",
// "topic": "2. Branching (Conditional Statements)",
// "difficulty": "hard",
// "context": "User Access Control",
// "title": "Validated Access Control System",
// "description": "Determine user access level using a switch statement considering: role, verification status, and login attempts. If the user has 3 or more login attempts, print 'Access Denied'. If the user is an admin and verified, print 'Full Access Granted'. If the user is an editor and verified, print 'Edit Access Granted'. If the user is a viewer and verified, print 'View Access Granted'. If the user has a known role but is not verified, print 'Please verify your account'. For any unknown role, print 'Unknown role'.",
// "answer_options": "case loginAttempts > 3:; case role == "admin" && isVerified:; case role == "viewer" && isVerified:; case role == "admin" || role == "editor" || role == "viewer":; case role = "admin" || "editor" || "viewer":; case loginAttempts int => 3"
package main

import "fmt"

func main() {
	var role string = "admin"
	var isVerified bool = true
	var loginAttempts int = 1
	switch {
	// TODO: Case statement to check if the user has more than 3 login attempts
	case loginAttempts > 3:
		fmt.Println("Access Denied")
	// TODO: Case statement to check if the user is an admin and verified
	case role == "admin" && isVerified:
		fmt.Println("Full Access Granted")
	case role == "editor" && isVerified:
		fmt.Println("Edit Access Granted")
	// TODO: Case statement to check if the user is a viewer and verified
	case role == "viewer" && isVerified:
		fmt.Println("View Access Granted")
	// TODO: Case statement to check if the user has a known role but is not verified
	case role == "admin" || role == "editor" || role == "viewer":
		fmt.Println("Please verify your account")
	default:
		fmt.Println("Unknown role")
	}
}
