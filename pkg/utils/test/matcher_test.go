package utils_test

import (
	"testing"
	"tiny-compiler/pkg/utils"

	"github.com/stretchr/testify/assert"
)

func Test_IsParen(t *testing.T) {
	// Test cases for IsParen function.
	tests := []struct {
		name     string
		input    rune
		expected bool
	}{
		{
			name:     "Parenthesis character '('",
			input:    '(',
			expected: true,
		},
		{
			name:     "Parenthesis character ')'",
			input:    ')',
			expected: true,
		},
		{
			name:     "Non-parenthesis character",
			input:    'a',
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := utils.IsParen(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func Test_IsNum(t *testing.T) {
	// Test cases for IsNum function.
	tests := []struct {
		name     string
		input    rune
		expected bool
	}{
		{
			name:     "Numeric digit '0'",
			input:    '0',
			expected: true,
		},
		{
			name:     "Numeric digit '5'",
			input:    '5',
			expected: true,
		},
		{
			name:     "Non-numeric character",
			input:    'a',
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := utils.IsNum(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func Test_IsLetter(t *testing.T) {
	// Test cases for IsLetter function.
	tests := []struct {
		name     string
		input    rune
		expected bool
	}{
		{
			name:     "Alphabetic character 'a'",
			input:    'a',
			expected: true,
		},
		{
			name:     "Alphabetic character 'Z'",
			input:    'Z',
			expected: true,
		},
		{
			name:     "Non-alphabetic character",
			input:    '1',
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := utils.IsLetter(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}
