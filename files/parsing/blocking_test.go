package parsing

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

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
