// "id": "pr_3_3_e",
// "type": "problem",
// "topic": "3. Looping",
// "difficulty": "hard",
// "context": "Data Grid",
// "title": "Matrix Cell Mapper",
// "description": "Iterate through rows and columns of a data grid. Print the grid symbol and coordinates using a nested for loop and Multi Variable Declarations. Declare string variables for grid symbol and cell marker; then declare integer variables for row and column limits using multi variable declaration. Use nested for loops to iterate through rows and columns, printing the grid symbol, cell marker, and current coordinates in each iteration.",
// "answer_options": "var rowLimit, colLimit int = 4, 3 | for r := 1; r <= rowLimit; r++ { | for c := 1; c <= colLimit; c++ { | fmt.Println(gridSymbol, cellMarker, r, c) | for c := 1; c <= rowLimit; c+1 { | for r := 1; r <= colLimit; r+1 {"
package main

import "fmt"

func main() {
	var gridSymbol, cellMarker string = "@", "Loc"
	// TODO: Declare integer variables 'rowLimit' and 'colLimit' using Multi Variable Declaration and assign them values
	var rowLimit, colLimit int = 4, 3
	// TODO: Create outer regular for loop to iterate 'r' from 1 to 'rowLimit' while incrementing 'r' by 1
	for r := 1; r <= rowLimit; r++ {
		// TODO: Create inner regular for loop to iterate 'c' from 1 to 'colLimit' while incrementing 'c' by 1
		for c := 1; c <= colLimit; c++ {
			// TODO: Print the values of 'gridSymbol', 'cellMarker', 'r', and 'c'
			fmt.Println(gridSymbol, cellMarker, r, c)
		}
	}
}