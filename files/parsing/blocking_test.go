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
					WordCount: 5,
					FileName:  base + "/r1.txt",
				},
				{
					WordCount: 2,
					FileName:  base + "/r2.txt",
				},
			},
		},
	}

	for _, test := range tests {
		result := BlockingFileStats(test.directory)
		assert.EqualValues(t, test.expected, result)

	}

}
