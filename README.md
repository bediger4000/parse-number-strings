# Daily Coding Problem: Problem #676 [Hard]

This problem was asked by LinkedIn.

Given a string, return whether it represents a number.
Here are the different kinds of numbers:

*   "10", a positive integer
*   "-10", a negative integer
*   "10.1", a positive real number
*   "-10.1", a negative real number
*   "1e5", a number in scientific notation

And here are examples of non-numbers:

*   "a"
*   "x 1"
*   "a -2"
*   "-"

## Analysis

The question as quoted above, seems really wide open.
Do you get to do this with standard libraries?
Can you use your programming language's parse-a-number feature,
whatever it may be?

* [Standard library conversion functions](v1.go)
* [Hand-coded state machine](v2.go)
* [Regular expression](v3.go)

### State Machine

The state machine I ended up with is more elaborate than I'd like.

![state machine](statemachine.png?raw=true)
