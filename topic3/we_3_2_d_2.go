// "id": "we_3_2_d_2",
// "type": "worked_example",
// "topic": "3. Looping",
// "difficulty": "medium",
// "context": "Device Testing",
// "title": "Device Error Check",
// "description": "Inspect device test cycles using a while-like loop. Note: conceptual similarity applied — uses a while-like loop with a condition instead of a standard 3-clause for loop, and explicit variable declarations."
package main

import "fmt"

func main() {
	// TODO: Declare a string variable 'errorType' and assign it a value
	var errorType string = "Failure"
	// TODO: Declare an integer variable 'checkTest' and assign it a value
	var checkTest int = 152
	// TODO: Declare an integer variable 'maxTests' and assign it a value
	var maxTests int = 154
	// TODO: Declare an integer variable 'testIndex' and assign it a value of 1
	var testIndex int = 1
	// TODO: Create a for loop with only a condition to check if 'testIndex' is less than or equal to 'maxTests'
	for testIndex <= maxTests {
		// TODO: Create an if statement to check if 'testIndex' is not equal to 'checkTest'
		if testIndex != checkTest {
			// Increment 'testIndex' manually to avoid an infinite loop before continuing
			testIndex++
			// TODO: Use 'continue' statement to skip the current iteration
			continue
		}
		// TODO: Print 'errorType' and the current test index 'testIndex'
		fmt.Println(errorType, testIndex)
		// TODO: Increment 'testIndex' manually to avoid an infinite loop
		testIndex++
	}
}