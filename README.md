# Kargo Summer Intern Pre-assessment 2021

This repository contains my solution to Kargo's Summer Intern Pre-assessment question.

Below is a description of the `phonetic` package, followed by the pre-assessment question.

## Phonetic Package
- `phonetic` contains the type alias `PhoneticDigits` of a 10-string array, each string representing the phonetic equivalent of a base-10 digit as seen in `digits`. An integer can be "split" into its digits, and these digits can be mapped to these strings by the indices of this array.
- `phoneticizeDigit` handles digits 0-9, and returns the corresponding string from `digits` and an error.
- `PhoneticizeInt` handles all positive integers, and use the strings returned from `PhoneticizeDigit` to create and return a string of phonetic digits joined without spaces, and returns an error.


### Question
"Convert an array of integers into an array of strings representing the phonetic equivalent of the
integer.

For example:
Given an array: [3, 25, 209], print “Three,TwoFive,TwoZeroNine” into stdout.  

Given an array: [10, 300, 5], print “OneZero,ThreeZeroZero,Five” into stdout."