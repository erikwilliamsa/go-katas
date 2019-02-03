package stringkata_test

import (
	"testing"

	"github.com/erikwilliamsa/go-katas/stringkata"
	"github.com/stretchr/testify/assert"
)

func TestChangePath(t *testing.T) {

	tests := []struct {
		current  string
		dest     string
		expected string
		message  string
	}{
		{
			current:  "/foo",
			dest:     "bar",
			expected: "bar",
			message:  "just bar",
		},
		{
			current:  "/foo/biz",
			dest:     "/foo/bar",
			expected: "../bar",
			message:  "foo biz in",
		},
		{
			current:  "/foo/biz/stuff",
			dest:     "/foo/bar",
			expected: "../../bar",
			message:  "2 up",
		},
		{
			current:  "/foo/biz/stuff",
			dest:     "/foo/bar/tacos/sales",
			expected: "../../bar/tacos/sales",
			message:  "long dest",
		},
		{
			current:  "/foo",
			dest:     "/foo/bar/tacos/sales",
			expected: "bar/tacos/sales",
			message:  "long dest, same  dir",
		},
	}

	for _, test := range tests {

		actual := stringkata.ChangePath(test.current, test.dest)
		assert.Equal(t, test.expected, actual, test.message)
	}
}
