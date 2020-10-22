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
whatever it may be? PHP will happily [let you do](x.php):

```php
$X = "12";
$N = 5 + $X;
```

assigning the number 17 to `$N`.

Python and Perl have similar convert-string-to-number features
that can probably be used in the context of this problem.

I did this problem in Go. I found 3 ways to do it.


* [Standard library conversion functions](v1.go),
uses Go's strconv functions to parse and verify.
* [Hand-coded state machine](v2.go) that only uses fmt, os packages.
* [Regular expression](v3.go), using a regular expression to
decided number/not a number

There's probably other, less savory ways.

The standard library conversion function version is probably
gives the answers you'd most like.
Both regular expression and the state machine version will accept
as number representations ".0" "0.0E0", "019"
and a few other things that are OK,
but are probably not ever written down as numbers by a human.
I don't know how important this is.

The approach of using a language's standard library functions
to parse and evaluate a string seems too easy for the "[Hard]" rating,
but it certainly slides in under the problem statement.
The candidate should probably ask for clarification on this issue.

Even if the interviewer wanted the candidate to write some kind
of recognizer (like my state machine) the "[Hard]" rating seems wrong.
I'd call a state machine "[Medium]".

An interviewer who uses this as a whiteboarding question
should prepare for candidates using hacks
(like use of `strconv` package)
or alternatives like using regular expressions.
If the interviewer wants to see programming language skills,
or program design skills,
this question needs re-focusing.
It's too loosely stated.

### State Machine

The state machine I ended up with is more elaborate than I'd like.
I did have the most fun with the state machine version.
I want to record the state machine for posterity.

![state machine](statemachine.png?raw=true)

I created the PNG image like this:

```sh
$ dot -Tpng -o statemachine.png statemachine.dot`
```

File `statemachine.dot` has [GraphViz](statemachine.dot) dot-language format
input for a directed graph describing the state machine.
Edges of the graph have annotations that are the input that
causes a state transition. "EOS" means "end of string".
