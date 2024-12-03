// ./internal/content/markdown_test.go
package content

import (
	"testing"
)

func TestMarkdownToHtml(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "Empty input",
			input:    "",
			expected: "",
		},
		{
			name:     "Single paragraph",
			input:    "This is a simple paragraph.",
			expected: "<p>This is a simple paragraph.</p>",
		},
		{
			name:     "Single heading",
			input:    "# Heading 1",
			expected: "<h1>Heading 1</h1>",
		},
		{
			name:     "Unordered list",
			input:    "- Item 1\n- Item 2\n- Item 3",
			expected: "<ul><li>Item 1</li><li>Item 2</li><li>Item 3</li></ul>",
		},
		{
			name:     "Ordered list",
			input:    "1. First\n2. Second\n3. Third",
			expected: "<ol><li>First</li><li>Second</li><li>Third</li></ol>",
		},
		{
			name:     "Mixed headings and paragraphs",
			input:    "# Heading 1\nThis is a paragraph.\n## Heading 2\nAnother paragraph.",
			expected: "<h1>Heading 1</h1><p>This is a paragraph.</p><h2>Heading 2</h2><p>Another paragraph.</p>",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output := MarkdownToHtml(tt.input)
			if output != tt.expected {
				t.Errorf("MarkdownToHtml() = %q, want %q", output, tt.expected)
			}
		})
	}
}
