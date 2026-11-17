// "id": "we_2_3_e_2",
// "type": "worked_example",
// "topic": "2. Branching (Conditional Statements)",
// "difficulty": "hard",
// "context": "System Authorization Gate",
// "title": "Validate Authorization Gate",
// "description": "Verify system security clearance level by validating accounts and key approvals using nested if-else structures."
package main

import "fmt"

func main() {
	// TODO: Declare three type inference variables for clearance role, approval key status, and revocation flag
	clearanceRole := "Manager"
	hasApprovalKey := true
	revocationFlag := false
	// TODO: Create an if statement to check if 'revocationFlag' is true
	if revocationFlag {
		fmt.Println("Access Blocked: Account Revoked")
	// TODO: Create an else statement to to handle false 'revocationFlag' and validate clearance levels based on role and key
	} else {
		// TODO: Create an if statement to validate clearanceRole is Manager and hasApprovalKey is true
		if clearanceRole == "Manager" && hasApprovalKey {
			fmt.Println("Full Clearance Confirmed")
		// TODO: Create an else if statement to check if the clearance role is Staff and has approval key
		} else if clearanceRole == "Staff" && hasApprovalKey {
			fmt.Println("Partial Clearance Confirmed")
		// TODO: Create an else if statement to check if the clearance role is Manager OR Staff but does not have approval key
		} else if clearanceRole == "Manager" || clearanceRole == "Staff" {
			fmt.Println("Pending Physical Approval Key")
		// TODO: Create an else statement to handle unrecognized clearance roles or missing approval keys
		} else {
			fmt.Println("Unregistered Profile Alert")
		}
	}
}