package parsing

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetTextFileStats(t *testing.T) {
	base := "./_resources/"

	tests := []struct {
		path     string
		expected TextFileStats
	}{
		{
			path: base + "/r1.txt",
			expected: TextFileStats{
				WordCount: 5,
				FileName:  base + "/r1.txt",
				CharCounts: map[string]int{
					"A": 1,
					"b": 1,
					"a": 2,
					"c": 2,
					"o": 4,
					"n": 1,
					"t": 1,
					"i": 1,
					"s": 1,
					"g": 1,
					"d": 1,
				},
			},
		},
		{
			path: base + "/r2.txt",
			expected: TextFileStats{
				WordCount: 2,
				FileName:  base + "/r2.txt",
				CharCounts: map[string]int{
					"f": 1,
					"o": 2,
					"a": 1,
					"b": 1,
					"r": 1,
				},
			},
		},
		{
			path: base + "/empty/empty.txt",
			expected: TextFileStats{
				WordCount:  0,
				FileName:   base + "/empty/empty.txt",
				CharCounts: map[string]int{},
			},
		},
	}

	for _, tt := range tests {
		file, err := os.Open(tt.path)
		assert.Nil(t, err, "Should be no file errors")

		result := GetTextFileStats(file)
		assert.Equal(t, tt.expected.WordCount, result.WordCount)
		assert.Equal(t, tt.expected.FileName, result.FileName, "File names should match")
		assert.Equal(t, tt.expected.CharCounts, result.CharCounts, "File names should match")
		file.Close()

	}
}

func TestBlockingAllFileStats(t *testing.T) {
	base := "./_resources"
	tests := []struct {
		directory string
		expected  DirTxtStats
	}{
		{
			directory: base,
			expected: DirTxtStats{
				TotalCount: 7,
				TotalCharCounts: map[string]int{
					"f": 1,
					"r": 1,
					"A": 1,
					"b": 2,
					"a": 3,
					"c": 2,
					"o": 6,
					"n": 1,
					"t": 1,
					"i": 1,
					"s": 1,
					"g": 1,
					"d": 1,
				},
				TextFileStats: []TextFileStats{
					TextFileStats{
						WordCount: 5,
						FileName:  base + "/r1.txt",
						CharCounts: map[string]int{
							"A": 1,
							"b": 1,
							"a": 2,
							"c": 2,
							"o": 4,
							"n": 1,
							"t": 1,
							"i": 1,
							"s": 1,
							"g": 1,
							"d": 1,
						},
					},
					TextFileStats{
						WordCount: 2,
						FileName:  base + "/r2.txt",
						CharCounts: map[string]int{
							"f": 1,
							"o": 2,
							"a": 1,
							"b": 1,
							"r": 1,
						},
					},
				},
			},
		},
	}

	for _, test := range tests {
		result := BlockingAllFileStats(test.directory)
		assert.Equal(t, test.expected.TotalCount, result.TotalCount)
		assert.Equal(t, test.expected.TotalCharCounts, result.TotalCharCounts)
		assert.Equal(t, test.expected.TextFileStats, result.TextFileStats)

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
