package main

/* Determine whether a string represents a number or not
 * based on regular expression matching.
 */

import (
	"fmt"
	"os"
	"regexp"
)

func main() {
	for _, stringNumber := range os.Args[1:] {
		if isNumber(stringNumber) {
			fmt.Printf("%q represents a number\n", stringNumber)
		} else {
			fmt.Printf("%q does not represent a number\n", stringNumber)
		}
	}
}

var numberPattern = regexp.MustCompile(`^[+-]*\.*[0-9][0-9]*\.*[0-9]*([eE][+-]*[0-9][0-9]*)*$`)

func isNumber(stringNumber string) bool {
	return numberPattern.Match([]byte(stringNumber))
}
