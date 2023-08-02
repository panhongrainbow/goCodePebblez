package bytez

import (
	"bytes"
	"strings"
	"testing"
)

// >>>>> >>>>> >>>>> >>>>> >>>>> Quotation

func Test_Check_Quotation(t *testing.T) {
	input := []byte("hello")
	expected := []byte(`"hello"`)
	result := Quotation(input)

	if !bytes.Equal(result, expected) {
		t.Errorf("Expected %s, but got %s", expected, result)
	}
}

func Benchmark_Performance_Quotation(b *testing.B) {
	input := []byte("benchmark")
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_ = Quotation(input)
	}
}

// >>>>> >>>>> >>>>> >>>>> >>>>> HasPrefix

func Test_Check_HasPrefix(t *testing.T) {
	testCases := []struct {
		a        []byte
		b        []byte
		expected bool
	}{
		{a: []byte("hello"), b: []byte("hello"), expected: true},
		{a: []byte("hello"), b: []byte("world"), expected: false},
		{a: []byte("prefix"), b: []byte("pre"), expected: true},
	}

	for _, tc := range testCases {
		result := HasPrefix(tc.a, tc.b)
		if result != tc.expected {
			t.Errorf("For input (%s, %s), expected %v but got %v", tc.a, tc.b, tc.expected, result)
		}
	}
}

func Benchmark_Performance_HasPrefix(b *testing.B) {
	a := []byte("prefix")
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_ = HasPrefix(a, []byte("pre"))
	}
}

// >>>>> >>>>> >>>>> >>>>> >>>>> StringToReadOnlyBytes

func Test_Check_StringToReadOnlyBytes(t *testing.T) {
	testCases := []struct {
		input    string
		expected []byte
	}{
		{input: "Hello", expected: []byte("Hello")},
		{input: "World", expected: []byte("World")},
		{input: "", expected: []byte("")},
	}

	for _, tc := range testCases {
		result := StringToReadOnlyBytes(tc.input)
		if !strings.EqualFold(string(result), tc.input) {
			t.Errorf("Expected %s, but got %s", tc.input, result)
		}
	}
}

func Benchmark_Performance_StringToReadOnlyBytes(b *testing.B) {
	input := "Hello World"
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_ = StringToReadOnlyBytes(input)
	}
}
