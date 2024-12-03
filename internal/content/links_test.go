// ./internal/content/links_test.go
package content

import (
	"testing"
)

func TestParseLinks(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "No links",
			input:    "This is a paragraph.",
			expected: "This is a paragraph.",
		},
		{
			name:     "Single link",
			input:    "Check out [Google](https://google.com).",
			expected: `Check out <a href="https://google.com">Google</a>.`,
		},
		{
			name:     "Multiple links",
			input:    "Visit [Google](https://google.com) and [OpenAI](https://openai.com).",
			expected: `Visit <a href="https://google.com">Google</a> and <a href="https://openai.com">OpenAI</a>.`,
		},
		{
			name:     "Link at start",
			input:    "[Google](https://google.com) is a search engine.",
			expected: `<a href="https://google.com">Google</a> is a search engine.`,
		},
		{
			name:     "Link at end",
			input:    "Search with [Google](https://google.com)",
			expected: `Search with <a href="https://google.com">Google</a>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output := parseLinks(tt.input)
			if output != tt.expected {
				t.Errorf("parseLinks() = %q, want %q", output, tt.expected)
			}
		})
	}
}
