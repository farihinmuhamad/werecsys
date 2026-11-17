// "id": "we_3_2_c_3",
// "type": "worked_example",
// "topic": "3. Looping",
// "difficulty": "medium",
// "context": "Data Recovery",
// "title": "Recovery Chunk Decrementer",
// "description": "Recover storage chunks in reverse order using a decrementing for loop. Note: deliberate differences applied — uses a 3-clause for loop, decrementing counter, and short variable declarations."
package main

import "fmt"

func main() {
	// TODO: Declare a type inference string variable 'chunkMarker' and assign it a value
	chunkMarker := "Chunk"
	// TODO: Declare a type inference integer variable 'startBlock' and assign it a value
	startBlock := 136
	// TODO: Declare a type inference integer variable 'endBlock' and assign it a value
	endBlock := 1
	// TODO: Create a standard 3-clause for loop starting from 'startBlock' down to 'endBlock' while decrementing the counter
	for b := startBlock; b >= endBlock; b-- {
		// TODO: Print the 'chunkMarker' and the current block number in each iteration
		fmt.Println(chunkMarker, b)
	}
}