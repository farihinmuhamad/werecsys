// "id": "we_2_3_d_1",
// "type": "worked_example",
// "topic": "2. Branching (Conditional Statements)",
// "difficulty": "hard",
// "context": "Spacecraft Launch Sequence",
// "title": "Launch Countdown Steps",
// "description": "Execute subsequent launch countdown operations using an integer switch statement with fallthrough."
package main

import "fmt"

func main() {
	// TODO: Declare type inference integer variable 'countdownStage' and assign it a value
	countdownStage := 3
	// TODO: Use switch statement to evaluate 'countdownStage' and print corresponding launch actions with fallthrough.
	switch countdownStage {
	// TODO: Case statement to check if 'countdownStage' is equal to 3
	case 3:
		fmt.Println("Ignition Sequence")
		// TODO: Use fallthrough to continue to the next case
		fallthrough
	// TODO: Case statement to check if 'countdownStage' is equal to 2
	case 2:
		fmt.Println("Engine Startup")
		fallthrough
	// TODO: Case statement to check if 'countdownStage' is equal to 1
	case 1:
		fmt.Println("Release Clamps")
	// TODO: Use default for determining Default Case for other countdown stages
	default:
		fmt.Println("Liftoff Confirmed")
	}
}
