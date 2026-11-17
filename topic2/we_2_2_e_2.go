// "id": "we_2_2_e_2",
// "type": "worked_example",
// "topic": "2. Branching (Conditional Statements)",
// "difficulty": "medium",
// "context": "Acoustic Noise Assessment",
// "title": "Check Environmental Noise Level",
// "description": "Assess noise sensor levels from two sensors using integer parameters and determine the noise status."
package main

import "fmt"

func main() {
	// TODO: Declare (multi) type inference integer variables 'noiseSensor1' and 'noiseSensor2' and assign them values
	noiseSensor1, noiseSensor2 := 82, 94
	// TODO: Use switch statement without an expression to evaluate the noise levels
	switch {
	// TODO: Case statement to check if 'noiseSensor1' is greater than 93 OR 'noiseSensor2' is greater than 93
	case noiseSensor1 > 93 || noiseSensor2 > 93:
		// TODO: Print the Deafening Noise level message
		fmt.Println("Deafening Noise")
	// TODO: Case statement to check if 'noiseSensor1' is greater than 77 OR 'noiseSensor2' is greater than 77
	case noiseSensor1 > 77 || noiseSensor2 > 77:
		// TODO: Print the Loud Noise level message
		fmt.Println("Loud Noise")
	// TODO: Case statement to check if 'noiseSensor1' is greater than 54 OR 'noiseSensor2' is greater than 54
	case noiseSensor1 > 54 || noiseSensor2 > 54:
		// TODO: Print the Moderate Noise level message
		fmt.Println("Moderate Noise")
	default:
		// TODO: Default case to print Quiet Noise status message
		fmt.Println("Quiet Noise")
	}
}
