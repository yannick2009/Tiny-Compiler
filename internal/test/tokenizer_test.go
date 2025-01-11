package internal_test

import (
	"testing"
	"tiny-compiler/internal"
	"tiny-compiler/pkg/enum"

	"github.com/stretchr/testify/assert"
)

func Test_Tokenizer(t *testing.T) {
	data := []struct {
		Name     string
		Input    string
		Expected []internal.Token
	}{
		{
			Name:     "Empty input",
			Input:    "",
			Expected: []internal.Token{},
		},
		{
			Name:  "Single parenthesis",
			Input: "(",
			Expected: []internal.Token{
				{Kind: enum.Paren, Value: "("},
			},
		},
		{
			Name:  "Single letter",
			Input: "a",
			Expected: []internal.Token{
				{Kind: enum.Letter, Value: "a"},
			},
		},
		{
			Name:  "Single number",
			Input: "1",
			Expected: []internal.Token{
				{Kind: enum.Number, Value: "1"},
			},
		},
		{
			Name:  "Multiple tokens",
			Input: "(a1)",
			Expected: []internal.Token{
				{Kind: enum.Paren, Value: "("},
				{Kind: enum.Letter, Value: "a"},
				{Kind: enum.Number, Value: "1"},
				{Kind: enum.Paren, Value: ")"},
			},
		},
		{
			Name:  "Complex input",
			Input: "(aa) (bbb)",
			Expected: []internal.Token{
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
