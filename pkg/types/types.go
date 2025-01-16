package types

import "tiny-compiler/pkg/enum"

// Node represents a single node in the Abstract Syntax Tree (AST).
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

// AST represents the Abstract Syntax Tree (AST) of the input tokens.
type AST Node

// Token represents a single token in the input string.
type Token struct {
	Kind  enum.TokenType
	Value string
}
