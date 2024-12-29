package enum

type TokenType int

const (
	Paren TokenType = iota
	Number
	Letter
)


// tokenTypeMap is a map that associates TokenType constants with their string representations.
// This map is used to convert token types to their corresponding string names for easier identification and debugging.
// The keys are of type TokenType and the values are their respective string names.
var tokenTypeMap map[TokenType]string = map[TokenType]string{
	Paren:         "Paren",
	Number:        "NumberLiteral",
	Letter:        "Name",
}

func (t TokenType) string() string {
	return tokenTypeMap[t]
}
