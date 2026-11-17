// "id": "we_1_3_d_2",
// "type": "worked_example",
// "topic": "1. Variables, Data Types, and Operators",
// "difficulty": "hard",
// "context": "Financial System",
// "title": "Loan Eligibility Verification",
// "description": "Verify client loan eligibility status using float64 comparisons, boolean flags, and combined logical evaluations."
package main

import "fmt"

func main() {
	// TODO: Declare a float64 variable 'loanAmount' and assign it a value
	loanAmount := 14800.75
	// TODO: Declare a float64 constant 'limitAmount' and assign it a value
	const limitAmount float64 = 20000.25
	// TODO: Declare a boolean variable 'hasCoSigner', initialize it with true
	var hasCoSigner bool = true
	// TODO: Declare a boolean variable 'isPreApproved', initialize it with false
	var isPreApproved bool = false
	// TODO: Determine limit safety status
	isLimitSafe := loanAmount <= limitAmount
	// TODO: Determine status checklist validation
	isStatusValid := hasCoSigner || isPreApproved
	// TODO: Verify if loan is approved combining both checks
	isLoanApproved := isLimitSafe && isStatusValid
	// TODO: Print the value of 'isLoanApproved'
	fmt.Println("Approval status:", isLoanApproved)
}