package types

import "tiny-compiler/pkg/enum"

// Node represents a single node in the Abstract Syntax Tree (AST). 
// It consists of a node type (Kind), a node value (Value), a node name (Name), and a list of child nodes (Body).
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

// Token represents a single token in the input string. It consists of a token type (Kind) and a token value (Value).
// The token type is defined in the TokenType enum, and the token value is the actual character or string that the token represents.
type Token struct {
	Kind  enum.TokenType
	Value string
}

// Visitor is a map of node types to visitor functions. It is used to traverse the AST and apply a function to each node.
// n is the current node being visited, and p is the parent node of n.
type Visitor map[string]func(n *Node, p Node)
