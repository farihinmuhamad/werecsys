// "id": "pr_3_1_a",
// "type": "problem",
// "topic": "3. Looping",
// "difficulty": "easy",
// "context": "Sequence Generation",
// "title": "Sum of First N Integers",
// "description": "Calculate the sum of the first N integers using a regular for loop and Regular Manifest Typing Variable Declaration. Declare a variable for the maximum number of steps and another for the total sum, then iterate through the sequence using a for loop to compute the sum.",
// "answer_options": "var maxSeqSteps int = 17 | var totalSum int = 0 | for i := 1; i <= maxSeqSteps; i++ { | totalSum += i | totalSum = totalSum + i++ | for i = 1 to i <= maxSeqSteps {"
package main

import "fmt"

func main() {
	// TODO: Declare an integer variable 'maxSeqSteps' and assign it a value
	var maxSeqSteps int = 17
	// TODO: Declare an integer variable 'totalSum' and assign it a value
	var totalSum int = 0
	// TODO: Create a regular for loop to iterate from 1 to 'maxSeqSteps' while incrementing 'i' by 1 at each iteration
	for i := 1; i <= maxSeqSteps; i++ {
		// TODO: Add the current index 'i' to 'totalSum'
		totalSum += i
	}
	fmt.Println(totalSum)
}