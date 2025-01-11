package internal_test

import (
	"log"
	"testing"
	"tiny-compiler/internal"

	"github.com/stretchr/testify/assert"
)

func TestParser(t *testing.T) {
	tokens := internal.Tokenizer("(add 2 (subtract 4 2))")
	expected := internal.AST{
		Kind: "Program",
		Body: []internal.Node{
			{
				Kind: "CallExpression",
				Name: "add",
				Params: []internal.Node{
					{
						Kind:  "NumberLiteral",
						Value: "2",
					},
					{
						Kind: "CallExpression",
						Name: "subtract",
						Params: []internal.Node{
							{
								Kind:  "NumberLiteral",
								Value: "4",
							},
							{
								Kind:  "NumberLiteral",
								Value: "2",
							},
						},
					},
				},
			},
		},
	}

	t.Run("Test Parser", func(t *testing.T) {
		ast := internal.Parser(tokens)
		log.Println(ast)
		assert.Equal(t, expected.Kind, ast.Kind, "AST kind mismatch")
		assert.Equal(t, len(expected.Body), len(ast.Body), "AST body length mismatch")

		for i, node := range ast.Body {
			assert.Equal(t, expected.Body[i].Kind, node.Kind, "Node kind mismatch")
			assert.Equal(t, expected.Body[i].Name, node.Name, "Node name mismatch")
			assert.Equal(t, expected.Body[i].Value, node.Value, "Node value mismatch")
			assert.Equal(t, len(expected.Body[i].Params), len(node.Params), "Node params length mismatch")

			for j, param := range node.Params {
				assert.Equal(t, expected.Body[i].Params[j].Kind, param.Kind, "Param kind mismatch")
				assert.Equal(t, expected.Body[i].Params[j].Value, param.Value, "Param value mismatch")
			}
		}
	})
}
