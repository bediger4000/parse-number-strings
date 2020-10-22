package main

/*
 * Decide if a string represents a number using Go standard
 * library functions.
 */

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	stringNumber := os.Args[1]

	if _, err := strconv.Atoi(stringNumber); err == nil {
		fmt.Printf("%q represents a number\n", stringNumber)
		return
	}

	if _, err := strconv.ParseFloat(stringNumber, 64); err == nil {
		fmt.Printf("%q represents a number\n", stringNumber)
		return
	}

	fmt.Printf("%q does not represent a number\n", stringNumber)
}
