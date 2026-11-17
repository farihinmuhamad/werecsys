// "id": "we_3_3_e_2",
// "type": "worked_example",
// "topic": "3. Looping",
// "difficulty": "hard",
// "context": "Sheet Analyzer",
// "title": "Table Cell Scanner",
// "description": "Iterate through rows and columns of a spreadsheet. Print the table symbol and cell coordinates using nested while-like loops and type inference variable declarations. Declare variables for table symbol, cell tag, row maximum, and column maximum. Note: conceptual similarity applied — uses nested while-like loops instead of standard 3-clause for loops, and type inference variable declarations.",

package main

import "fmt"

func main() {
	tableSym := "$"
	cellTag := "Tab"
	rowMax, colMax := 188, 189
	r := 1
	// TODO: Create outer loop with only a condition to check if 'r' is less than or equal to 'rowMax'
	for r <= rowMax {
		c := 1
		// TODO: Create inner loop with only a condition to check if 'c' is less than or equal to 'colMax'
		for c <= colMax {
			// TODO: Print the values of 'tableSym', 'cellTag', 'r', and 'c'
			fmt.Println(tableSym, cellTag, r, c)
			// TODO: Increment 'c' manually inside the inner loop
			c++
		}
		// TODO: Increment 'r' manually inside the outer loop
		r++
	}
}