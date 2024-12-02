package content

import "testing"

func TestMarkdownToHtml(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "Full conversion with multiple syntax",
			input:    "# Heading **bold** _italic_",
			expected: "<h1>Heading <strong>bold</strong> <em>italic</em></h1>",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := MarkdownToHtml(tt.input)
			if result != tt.expected {
				t.Errorf("got %q, want %q", result, tt.expected)
			}
		})
	}
}

func TestParseHeadings(t *testing.T) {
	input := "# Heading 1"
	expected := "<h1>Heading 1</h1>"
	if result := parseHeadings(input); result != expected {
		t.Errorf("parseHeadings() = %q, want %q", result, expected)
	}
}

func TestParseBold(t *testing.T) {
	input := "**bold**"
	expected := "<strong>bold</strong>"
	if result := parseBold(input); result != expected {
		t.Errorf("parseBold() = %q, want %q", result, expected)
	}
}

func TestParseItalic(t *testing.T) {
	input := "_italic_"
	expected := "<em>italic</em>"
	if result := parseItalic(input); result != expected {
		t.Errorf("parseItalic() = %q, want %q", result, expected)
	}
}

func TestParseLinks(t *testing.T) {
	input := "[example](https://example.com)"
	expected := `<a href="https://example.com">example</a>`
	if result := parseLinks(input); result != expected {
		t.Errorf("parseLinks() = %q, want %q", result, expected)
	}
}

func TestParseUnorderedList(t *testing.T) {
	input := "- Item 1\n- Item 2"
	expected := "<ul><li>Item 1</li><li>Item 2</li></ul>"
	if result := parseUnorderedList(input); result != expected {
		t.Errorf("parseUnorderedList() = %q, want %q", result, expected)
	}
}

func TestParseOrderedList(t *testing.T) {
	input := "1. Item 1\n2. Item 2"
	expected := "<ol><li>Item 1</li><li>Item 2</li></ol>"
	if result := parseOrderedList(input); result != expected {
		t.Errorf("parseOrderedList() = %q, want %q", result, expected)
	}
}
