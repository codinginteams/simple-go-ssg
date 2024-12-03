// ./internal/content/headings_test.go
package content

import (
	"testing"
)

func TestParseHeadings(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "No headings",
			input:    "This is a paragraph.",
			expected: "This is a paragraph.",
		},
		{
			name:     "Heading level 1",
			input:    "# Heading 1",
			expected: "<h1>Heading 1</h1>",
		},
		{
			name:     "Heading level 3",
			input:    "### Heading 3",
			expected: "<h3>Heading 3</h3>",
		},
		{
			name:     "Multiple headings",
			input:    "# Heading 1\n## Heading 2\n### Heading 3",
			expected: "<h1>Heading 1</h1>\n<h2>Heading 2</h2>\n<h3>Heading 3</h3>",
		},
		{
			name:     "Heading with trailing spaces",
			input:    "#### Heading 4   ",
			expected: "<h4>Heading 4   </h4>",
		},
		{
			name:     "Heading with special characters",
			input:    "##### Heading @#$%^&*()!",
			expected: "<h5>Heading @#$%^&*()!</h5>",
		},
		{
			name:     "Non-heading lines mixed",
			input:    "# Heading 1\nThis is a paragraph.\n## Heading 2",
			expected: "<h1>Heading 1</h1>\nThis is a paragraph.\n<h2>Heading 2</h2>",
		},
		{
			name:     "Heading with multiple spaces",
			input:    "##    Heading with multiple spaces",
			expected: "<h2>   Heading with multiple spaces</h2>",
		},
		{
			name:     "Invalid headings",
			input:    "####Heading without space",
			expected: "####Heading without space",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output := parseHeadings(tt.input)
			if output != tt.expected {
				t.Errorf("parseHeadings() = %q, want %q", output, tt.expected)
			}
		})
	}
}
