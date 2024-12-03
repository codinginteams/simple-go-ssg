// ./internal/content/text_formats_test.go
package content

import (
	"testing"
)

func TestParseBold(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "Single bold",
			input:    "This is **bold** text.",
			expected: "This is <strong>bold</strong> text.",
		},
		{
			name:     "Bold at start",
			input:    "**Bold** is at the start.",
			expected: "<strong>Bold</strong> is at the start.",
		},
		{
			name:     "Bold at end",
			input:    "End with **bold**",
			expected: "End with <strong>bold</strong>",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output := parseBold(tt.input)
			if output != tt.expected {
				t.Errorf("parseBold(%q) = %q, want %q", tt.input, output, tt.expected)
			}
		})
	}
}

func TestParseItalic(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "Single italic",
			input:    "This is _italic_ text.",
			expected: "This is <em>italic</em> text.",
		},
		{
			name:     "Italic at start",
			input:    "_Italic_ is at the start.",
			expected: "<em>Italic</em> is at the start.",
		},
		{
			name:     "Italic at end",
			input:    "End with _italic_",
			expected: "End with <em>italic</em>",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output := parseItalic(tt.input)
			if output != tt.expected {
				t.Errorf("parseItalic(%q) = %q, want %q", tt.input, output, tt.expected)
			}
		})
	}
}
