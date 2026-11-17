// "id": "pr_2_1_a",
// "type": "problem",
// "topic": "2. Branching (Conditional Statements)",
// "difficulty": "easy",
// "context": "Pass or Fail Determination",
// "title": "Student Grading System",
// "description": "Determine if a student has passed or failed based on their score compared to a passing score. Declare a constant for the passing score and a variable for the student's score. Use an if statement to check if the student's score is greater than or equal to the passing score. If it is, print 'Pass'. Otherwise, print 'Fail'.",
// "answer_options": "const passingScore int = 70; var studentScore int = 75; if studentScore >= passingScore {; fmt.Println("Pass"); var passingScore float64 = 70.0; const studentScore int = 75"
package main

import "fmt"

func main() {
	// TODO: Declare a constant 'passingScore' of type int and assign it a value
	const passingScore int = 70
	// TODO: Declare a variable 'studentScore' of type int and assign it a value
	var studentScore int = 75
	// TODO: Create an if statement to check if 'studentScore' is greater than or equal to 'passingScore'
	if studentScore >= passingScore {
		// TODO: Inside the if block, print the value of "Pass"
		fmt.Println("Pass")
	}
}