package sanitize

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestInteger tests the Integer sanitize method
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

// BenchmarkInteger benchmarks the Integer method
func BenchmarkInteger(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Integer("a-bc12.3def987asdf--")
	}
}

// TestFloat tests the Float sanitize method
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

// BenchmarkFloat benchmarks the Float method
func BenchmarkFloat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Float("a-bc12.3def987asdf--")
	}
}

// TestIP tests the IP sanitize method
func TestIP(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		input    string
		expected string
	}{
		{"192.168.0.1", "192.168.0.1"},
		{"192.168.0.256", ""},
		{"192.168", ""},
		{"IP: 192.168.0.1", ""},
		{"2001:0db8:85a3:0000:0000:8a2e:0370:7334", "2001:db8:85a3::8a2e:370:7334"},
		{"2001:db8:0:0:1:0:0:1", "2001:db8::1:0:0:1"},
		{"d8c4:e12f:eff7:df64:3995:df25:aaca:49eb", "d8c4:e12f:eff7:df64:3995:df25:aaca:49eb"},
		{"6b71:3208:07ad:0629:a150:5734:15e8:950d", "6b71:3208:7ad:629:a150:5734:15e8:950d"},
		{"::ffff:c000:0280", "192.0.2.128"},
	}

	for _, test := range tests {
		output := IP(test.input)
		assert.Equal(t, test.expected, output)
	}
}

// BenchmarkIP benchmarks the IP method
func BenchmarkIP(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = IP("192.168.0.1")
	}
}
