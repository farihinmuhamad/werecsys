// "id": "we_3_2_d_3",
// "type": "worked_example",
// "topic": "3. Looping",
// "difficulty": "medium",
// "context": "Task Processing",
// "title": "Task Priority Checker",
// "description": "Check task priorities using a decrementing for loop. Note: deliberate differences applied — uses a decrementing loop and checks equality directly instead of using continue."
package main

import "fmt"

func main() {
	// TODO: Declare a type inference variable 'taskStatus' (string) and assign it a value
	taskStatus := "Fatal"
	// TODO: Declare a type inference variable 'targetTask' (int) and assign it a value
	targetTask := 164
	// TODO: Declare a type inference variable 'maxTasks' (int) and assign it a value
	maxTasks := 166
	// TODO: Create a decrementing for loop starting from 'maxTasks' down to 1 while decrementing 'taskIndex' by 1 at each iteration
	for taskIndex := maxTasks; taskIndex >= 1; taskIndex-- {
		// TODO: Create an if statement to check if 'taskIndex' is equal to 'targetTask'
		if taskIndex == targetTask {
			// TODO: Print 'taskStatus' and the current task index 'taskIndex'
			fmt.Println(taskStatus, taskIndex)
		}
	}
}