package sanitize

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestInteger tests the integer sanitize method
func TestInteger(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		input    string
		expected string
	}{
		{"45sDa8f$sd541zfa", "458541"},
		{"a-bc12.3def987asdf--", "-123987"},
		{"z,2134-59", "213459"},
		{".,-", ""},
	}

	for _, test := range tests {
		output := Integer(test.input)
		assert.Equal(t, test.expected, output)
	}
}

// BenchmarkInteger benchmarks the integer method
func BenchmarkInteger(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Integer("a-bc12.3def987asdf--")
	}
}

// TestFloat tests the float sanitize method
func TestFloat(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		input    string
		expected string
	}{
		{"45sDa8f$sd541zfa", "458541"},
		{"a-bc12.3def987asdf--", "-12.3987"},
		{"z,2134-59", "0,213459"},
		{".,-", ""},
	}

	for _, test := range tests {
		output := Float(test.input)
		assert.Equal(t, test.expected, output)
	}
}

// BenchmarkFloat benchmarks the float method
func BenchmarkFloat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Float("a-bc12.3def987asdf--")
	}
}
