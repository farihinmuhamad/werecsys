// "id": "we_3_3_c_3",
// "type": "worked_example",
// "topic": "3. Looping",
// "difficulty": "hard",
// "context": "Frame Analysis",
// "title": "Frame Decoder",
// "description": "Search a series of video frames for a target frame starting from the end. Print a decoding message, loop through iterations in reverse, skip trailing frames with continue, and stop when the target frame is found using break. Note: deliberate differences applied — uses a decrementing loop and short variable declarations.",

package main

import "fmt"

func main() {
	// TODO: Declare a type inference variable 'decodeMode' (string) and assign it a value
	decodeMode := "Extract"
	// TODO: Declare two type inference variables 'totalFrames' (int) and 'targetFrame' (int) and assign them values
	totalFrames, targetFrame := 174, 172
	// TODO: Create a decrementing for loop to iterate 'i' from 'totalFrames' down to 1
	for i := totalFrames; i >= 1; i-- {
		// TODO: Create an if statement to check if 'i' is greater than 'targetFrame'
		if i > targetFrame {
			// TODO: Execute 'continue' statement to skip frames greater than 'targetFrame'
			continue
		}
		fmt.Println(decodeMode, i)
		// TODO: Create an if statement to check if 'i' is equal to 'targetFrame'
		if i == targetFrame {
			// TODO: Execute 'break' statement to exit the loop when 'i' is equal to 'targetFrame'
			break
		}
	}
}