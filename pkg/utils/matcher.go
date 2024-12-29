package utils

import (
	"regexp"
	"tiny-compiler/pkg/constants"
)

/*
IsNum checks if the given rune represents a numeric digit (0-9).
It returns true if the rune is a digit, and false otherwise.

Parameters:
- char: The rune to be checked.

Returns:
- bool: true if the rune is a numeric digit, false otherwise.
*/
func IsNum(char rune) bool {
	if char == ' ' {
		return false
	}
	regex := regexp.MustCompile(constants.NumRegex)
	return regex.MatchString(string(char))
}

/*
IsLetter checks if the given rune represents an alphabetic character (a-z, A-Z).
It returns true if the rune is an alphabetic character, and false otherwise.

Parameters:
- char: The rune to be checked.

Returns:
- bool: true if the rune is an alphabetic character, false otherwise.
*/
func IsLetter(char rune) bool {
	if char == ' ' {
		return false
	}
	regex := regexp.MustCompile(constants.CharRegex)
	return regex.MatchString(string(char))
}


// IsParen checks if the given rune is a parenthesis character.
// It returns true if the rune is either '(' or ')', otherwise it returns false.
func IsParen(char rune) bool {
	if char == '(' || char == ')' {
		return true
	}
	return false
}
