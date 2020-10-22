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
	for _, stringNumber := range os.Args[1:] {
		if _, err := strconv.Atoi(stringNumber); err == nil {
			fmt.Printf("%q represents a number\n", stringNumber)
			continue
		}

		if _, err := strconv.ParseFloat(stringNumber, 64); err == nil {
			fmt.Printf("%q represents a number\n", stringNumber)
			continue
		}

		fmt.Printf("%q does not represent a number\n", stringNumber)
	}
}
