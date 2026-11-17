// "id": "we_3_3_e_1",
// "type": "worked_example",
// "topic": "3. Looping",
// "difficulty": "hard",
// "context": "Screen Renderer",
// "title": "Pixel Coordinate Mapper",
// "description": "Iterate through rows and columns of a screen display. Print the pixel character and coordinates using a nested for loop and Multi Variable Declarations. Declare string variables for pixel character and coordinate label; then declare integer variables for row and column limits using multi variable declaration. Use nested for loops to iterate through rows and columns, printing the pixel character, coordinate label, and current coordinates in each iteration.",

package main

import "fmt"

func main() {
	// TODO: Declare string variables 'pixelChar' and 'coordLabel' using Multi Variable Declaration and assign them values
	var pixelChar, coordLabel string = "?", "Pix"
	// TODO: Declare integer variables 'maxRows' and 'maxCols' using Multi Variable Declaration and assign them values
	var maxRows, maxCols int = 183, 187
	// TODO: Create outer regular for loop to iterate 'r' from 1 to 'maxRows' while incrementing 'r' by 1
	for r := 1; r <= maxRows; r++ {
		// TODO: Create inner regular for loop to iterate 'c' from 1 to 'maxCols' while incrementing 'c' by 1
		for c := 1; c <= maxCols; c++ {
			// TODO: Print the values of 'pixelChar', 'coordLabel', 'r', and 'c'
			fmt.Println(pixelChar, coordLabel, r, c)
		}
	}
}