package main

/*
 * Recognize strings that represent numbers via
 * artisanlly-crafted state machine.
 */

import (
	"fmt"
	"os"
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

	for _, stringNumber := range os.Args[1:] {

		state := isNumber(stringNumber)

		if state == success {
			fmt.Printf("%q represents a number\n", stringNumber)
		}
		if state == fail {
			fmt.Printf("%q does not represent a number\n", stringNumber)
		}
	}
}

func isNumber(stringNumber string) State {
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
			if isDigit(runes[idx]) {
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
			if isDigit(runes[idx]) {
				state = digit
			}
			idx++
		case dotfound:
			state = fail
			if idx >= ln {
				continue
			}
			if isDigit(runes[idx]) {
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
			if isDigit(runes[idx]) {
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
			if isDigit(runes[idx]) {
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
			if isDigit(runes[idx]) {
				state = expdigits
			}
			idx++
		case expsign:
			state = fail
			if idx >= ln {
				continue
			}
			if isDigit(runes[idx]) {
				state = expdigits
			}
			idx++
		case expdigits:
			state = fail
			if idx >= ln {
				state = success
				continue
			}
			if isDigit(runes[idx]) {
				state = expdigits
			}
			idx++
		case moredigits:
			state = fail
			if idx >= ln {
				state = success
				continue
			}
			if isDigit(runes[idx]) {
				state = moredigits
			}
			if runes[idx] == 'E' || runes[idx] == 'e' {
				state = foundexp
			}
			idx++
		}
	}
	return state
}

func isDigit(r rune) bool {
	if r >= '0' && r <= '9' {
		return true
	}
	return false
}
