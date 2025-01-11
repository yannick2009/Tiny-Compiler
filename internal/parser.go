package internal

import (
	"log"
	"tiny-compiler/pkg/enum"
)

type Node struct {
	Kind       string
	Value      string
	Name       string
	Callee     *Node
	Expression *Node
	Body       []Node
	Params     []Node
	Arguments  []*Node
	Context    []*Node
}

type AST Node

func NewATS() AST {
	return AST{
		Kind: "Program",
		Body: []Node{},
	}
}

// Parser is a function that takes a slice of Token structs and parses it into an Abstract Syntax Tree (AST).
// The function returns an AST struct, which represents the structure of the input tokens.
// The function iterates over the input tokens, combining them into higher-level structures based on their relationships.
// The function uses a recursive descent parsing algorithm to build the AST, starting from the top-level structure and working down to the individual tokens.
// The function returns the final AST once it has processed all the input tokens.
var cursor int
var tokens []Token

func Parser(inputTokens []Token) AST {
	tokens = inputTokens
	ATS := NewATS()

	for cursor < len(tokens) {
		ATS.Body = append(ATS.Body, walk())
	}

	return ATS
}

// walk traverses through tokens and returns a Node structure based on the current token.
// It processes parentheses and number tokens, advancing the cursor position after each token is processed.
// Returns an empty Node if the current token doesn't match any expected patterns.
func walk() Node {
	currentNode := tokens[cursor]

	if currentNode.Kind == enum.Number {
		cursor++
		return Node{
			Kind:  enum.Number.String(),
			Value: currentNode.Value,
		}
	}

	if currentNode.Kind == enum.Paren && currentNode.Value == "(" {
		cursor++
		token := tokens[cursor]

		node := Node{
			Kind:   token.Kind.String(),
			Name:   token.Value,
			Params: []Node{},
		}

		cursor++
		token = tokens[cursor]

		for token.Kind != enum.Paren || (token.Kind == enum.Paren && token.Value != ")") {
			node.Params = append(node.Params, walk())
			token = tokens[cursor]
		}

		cursor++
		return node
	}

	// check if the current node is a letter
	log.Fatal(tokens[cursor].Kind.String())
	return Node{}
}
