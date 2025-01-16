package internal_test

import (
	"testing"
	"tiny-compiler/internal"
	"tiny-compiler/pkg/enum"
	"tiny-compiler/pkg/types"

	"github.com/stretchr/testify/assert"
)

func Test_Tokenizer(t *testing.T) {
	data := []struct {
		Name     string
		Input    string
		Expected []types.Token
	}{
		{
			Name:     "Empty input",
			Input:    "",
			Expected: []types.Token{},
		},
		{
			Name:  "Single parenthesis",
			Input: "(",
			Expected: []types.Token{
				{Kind: enum.Paren, Value: "("},
			},
		},
		{
			Name:  "Single letter",
			Input: "a",
			Expected: []types.Token{
				{Kind: enum.Letter, Value: "a"},
			},
		},
		{
			Name:  "Single number",
			Input: "1",
			Expected: []types.Token{
				{Kind: enum.Number, Value: "1"},
			},
		},
		{
			Name:  "Multiple tokens",
			Input: "(a1)",
			Expected: []types.Token{
				{Kind: enum.Paren, Value: "("},
				{Kind: enum.Letter, Value: "a"},
				{Kind: enum.Number, Value: "1"},
				{Kind: enum.Paren, Value: ")"},
			},
		},
		{
			Name:  "Complex input",
			Input: "(aa) (bbb)",
			Expected: []types.Token{
				{Kind: enum.Paren, Value: "("},
				{Kind: enum.Letter, Value: "aa"},
				{Kind: enum.Paren, Value: ")"},
				{Kind: enum.Paren, Value: "("},
				{Kind: enum.Letter, Value: "bbb"},
				{Kind: enum.Paren, Value: ")"},
			},
		},
	}

	for _, d := range data {
		t.Run(d.Name, func(t *testing.T) {
			t.Parallel()
			tokens := internal.Tokenizer(d.Input)
			if len(tokens) != len(d.Expected) {
				t.Errorf("Expected %d tokens, got %d", len(d.Expected), len(tokens))
			}

			assert.Equal(t, d.Expected, tokens)
		})
	}
}
