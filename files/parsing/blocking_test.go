package parsing

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBlockingFileStats(t *testing.T) {
	base := "./_resources"
	tests := []struct {
		directory string
		expected  []TextFileStats
	}{
		{
			directory: base,
			expected: []TextFileStats{
				{
					WordCount:  5,
					FileName:   base + "/r1.txt",
					CharCounts: map[string]int{},
				},
				{
					WordCount:  2,
					FileName:   base + "/r2.txt",
					CharCounts: map[string]int{},
				},
			},
		},
	}

	for _, test := range tests {
		result := BlockingFileStats(test.directory)
		assert.EqualValues(t, test.expected, result)

	}

}

func TestAlphaNumericCount(t *testing.T) {
	tests := []struct {
		input    string
		expected map[string]int
	}{
		{
			input: "AaA Bc",
			expected: map[string]int{
				"A": 2,
				"a": 1,
				"B": 1,
				"c": 1,
			},
		},
		{
			input: `AaA Bc 
					lmnoP QQ`,
			expected: map[string]int{
				"A": 2,
				"a": 1,
				"B": 1,
				"c": 1,
				"l": 1,
				"m": 1,
				"n": 1,
				"o": 1,
				"P": 1,
				"Q": 2,
			},
		},
		{
			input: `*&^A*D`,
			expected: map[string]int{
				"A": 1,
				"D": 1,
			},
		},
		{
			input:    ``,
			expected: map[string]int{},
		},
	}

	for _, test := range tests {
		assert.EqualValues(t,
			test.expected,
			AlphaNumericCount(test.input))

	}
}

func TestCountWords(t *testing.T) {

	tests := []struct {
		args     string
		expected int
	}{
		{
			args:     "a quick brown dog",
			expected: 4,
		},
		{
			args:     "{ a  *@ quick brown dog",
			expected: 4,
		},
		{
			args:     "12 34",
			expected: 2,
		},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.expected, CountWords(tt.args))
	}
}
