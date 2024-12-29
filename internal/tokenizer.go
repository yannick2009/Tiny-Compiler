package internal

import (
	"tiny-compiler/pkg/enum"
	"tiny-compiler/pkg/utils"
)

type Token struct {
	Kind  enum.TokenType
	Value string
}

func newToken(kind enum.TokenType, value string) Token {
	return Token{Kind: kind, Value: value}
}

/*
Tokenizer is a function that takes an input string and tokenizes it.
It returns a slice of Token structs, where each Token represents a single token in the input string.
The function iterates over the input string, character by character, and classifies each character into a token type.
The token types are defined in the TokenType enum, and the function creates a Token struct for each character based on its type.
The function then appends the Token struct to the Tokens slice and continues to the next character.
The function returns the Tokens slice once it has processed the entire input string.
*/
func Tokenizer(input string) []Token {
	if input == "" {
		return []Token{}
	}
	
	input += "\n"
	var Tokens []Token
	inputArr := []rune(input)
	currentPos := 0

	for currentPos < len(inputArr) {
		currentChar := inputArr[currentPos]
		// checking if the char is empty
		if currentChar == ' ' {
			currentPos++
			continue
		}

		// checking if it's a parenthese
		if utils.IsParen(currentChar) {
			Tokens = append(Tokens, newToken(enum.Paren, string(currentChar)))
			currentPos++
			continue
		}

		// checking if it's a name
		if utils.IsLetter(currentChar) {
			value := ""
			for utils.IsLetter(inputArr[currentPos]) {
				value += string(inputArr[currentPos])
				currentPos++
			}

			Tokens = append(Tokens, newToken(enum.Letter, value))
			continue
		}

		// checking if it's a number
		if utils.IsNum(currentChar) {
			value := ""
			for utils.IsNum(inputArr[currentPos]) {
				value += string(inputArr[currentPos])
				currentPos++
			}

			Tokens = append(Tokens, newToken(enum.Number, value))
			continue
		}

		if currentChar == '\n' {
			break
		}
	}

	return Tokens
}
