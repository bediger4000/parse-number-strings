package main

/*
 * Recognize strings that represent numbers via
 * artisanlly-crafted state machine.
 */

import (
	"fmt"
	"os"
	"unicode"
)

type State int

const (
	start        State = iota
	signfound    State = iota
	dotfound     State = iota
	digit        State = iota
	decimalpoint State = iota
	foundexp     State = iota
	moredigits   State = iota
	expsign      State = iota
	expdigits    State = iota
	success      State = iota
	fail         State = iota
)

func main() {
	stringNumber := os.Args[1]

	/* String representation of a number:
	* optional "-" or "+"
	* followed by one or more digits ("0" through "9")
	* optional "."
	* optional "e" or "E"
	  * followed by 1 non-zero digit, with an option digit after it
	*/

	var state State = start

	runes := []rune(stringNumber)
	idx := 0
	ln := len(runes)

	for state != fail && state != success {
		switch state {
		case start:
			state = fail
			if idx >= ln {
				continue
			}
			if runes[idx] == '+' || runes[idx] == '-' {
				state = signfound
			}
			if runes[idx] == '.' {
				state = dotfound
			}
			if unicode.IsDigit(runes[idx]) {
				state = digit
			}
			idx++
		case signfound:
			state = fail
			if idx >= ln {
				continue
			}
			if runes[idx] == '.' {
				state = dotfound
			}
			if unicode.IsDigit(runes[idx]) {
				state = digit
			}
			idx++
		case dotfound:
			state = fail
			if idx >= ln {
				continue
			}
			if unicode.IsDigit(runes[idx]) {
				state = digit
			}
			idx++
		case digit:
			state = fail
			if idx >= ln {
				state = success
				continue
			}
			if runes[idx] == '.' {
				state = decimalpoint
			}
			if unicode.IsDigit(runes[idx]) {
				state = digit
			}
			if runes[idx] == 'E' || runes[idx] == 'e' {
				state = foundexp
			}
			idx++
		case decimalpoint:
			state = fail
			if idx >= ln {
				state = success
				continue
			}
			if runes[idx] == 'E' || runes[idx] == 'e' {
				state = foundexp
			}
			if unicode.IsDigit(runes[idx]) {
				state = moredigits
			}
			idx++
		case foundexp:
			state = fail
			if idx >= ln {
				continue
			}
			if runes[idx] == '+' || runes[idx] == '-' {
				state = expsign
			}
			if unicode.IsDigit(runes[idx]) {
				state = expdigits
			}
			idx++
		case expsign:
			state = fail
			if idx >= ln {
				continue
			}
			if unicode.IsDigit(runes[idx]) {
				state = expdigits
			}
			idx++
		case expdigits:
			state = fail
			if idx >= ln {
				state = success
				continue
			}
			if unicode.IsDigit(runes[idx]) {
				state = expdigits
			}
			idx++
		case moredigits:
			state = fail
			if idx >= ln {
				state = success
				continue
			}
			if unicode.IsDigit(runes[idx]) {
				state = moredigits
			}
			if runes[idx] == 'E' || runes[idx] == 'e' {
				state = foundexp
			}
			idx++
		}
	}

	if state == success {
		fmt.Printf("%q represents a number\n", stringNumber)
	}
	if state == fail {
		fmt.Printf("%q does not represent a number\n", stringNumber)
	}
}
