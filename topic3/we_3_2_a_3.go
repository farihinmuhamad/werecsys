// "id": "we_3_2_a_3",
// "type": "worked_example",
// "topic": "3. Looping",
// "difficulty": "medium",
// "context": "Probe Calibration",
// "title": "Calibration Offset Reverse Accumulator",
// "description": "Accumulate probe offset values over multiple calibration runs using a decrementing for loop and short variable declarations. Note: deliberate differences applied — reverse decrement loop and short variable declarations."
package main

import "fmt"

func main() {
	// TODO: Declare a float64 variable 'probeOffset' and assign it a value using short declaration
	probeOffset := 0.44
	// TODO: Declare an integer variable 'calibRunCount' and assign it a value using short declaration
	calibRunCount := 44
	// TODO: Declare a float64 variable 'cumulativeOffset' and assign it a value using short declaration
	cumulativeOffset := 0.0
	// TODO: Create a regular for loop to iterate 'a' from 'calibRunCount' down to 1 while decrementing 'a' by 1
	for a := calibRunCount; a >= 1; a-- {
		// TODO: Change the value of 'cumulativeOffset' by adding 'probeOffset' to it inside the loop
		cumulativeOffset += probeOffset
	}
	fmt.Println(cumulativeOffset)
}