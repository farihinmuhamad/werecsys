// "id": "pr_3_2_d",
// "type": "problem",
// "topic": "3. Looping",
// "difficulty": "medium",
// "context": "Cargo Shipping",
// "title": "Cargo Weight Checker",
// "description": "Check cargo container weights, print warning message for specific indices using continue only without break, using Type Inference Variable Declarations. Declare variables for warning message, index to check, and maximum containers. Iterate through container indices, print warning for overweight containers at the check index, and skip others.",
// "answer_options": "for i := 1; i <= maxContainers; i++ { | if i != checkIndex { | continue | fmt.Println(warningMsg, i) | for i = 1; i <= maxContainers do { | if i <> checkIndex {"
package main

import "fmt"

func main() {
	warningMsg := "Overweight"
	checkIndex := 49
	maxContainers := 59
	// TODO: Create a regular for loop to iterate 'i' from 1 to 'maxContainers' with increment of 1 at each iteration
	for i := 1; i <= maxContainers; i++ {
		// TODO: Create an if statement to check if 'i' is not equal to 'checkIndex'
		if i != checkIndex {
			// TODO: Use 'continue' statement to skip the current iteration to the next iteration of the loop
			continue
		}
		// TODO: Print 'warningMsg' and the current index 'i' when 'i' is equal to 'checkIndex'
		fmt.Println(warningMsg, i)
	}
}