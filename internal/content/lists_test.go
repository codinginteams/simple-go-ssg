// ./internal/content/lists_test.go
package content

import (
	"testing"
)

func TestIsUnorderedList(t *testing.T) {
	tests := []struct {
		name     string
		line     string
		expected bool
	}{
		{"Unordered list item with single dash", "- Item 1", true},
		{"Ordered list item", "1. Item 1", false},
		{"Non-list line", "This is a paragraph.", false},
		{"Empty line", "", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := isUnorderedList(tt.line)
			if result != tt.expected {
				t.Errorf("isUnorderedList(%q) = %v, want %v", tt.line, result, tt.expected)
			}
		})
	}
}

func TestIsOrderedList(t *testing.T) {
	tests := []struct {
		name     string
		line     string
		expected bool
	}{
		{"Ordered list item with number", "1. Item 1", true},
		{"Unordered list item", "- Item 1", false},
		{"Non-list line", "This is a paragraph.", false},
		{"Empty line", "", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := isOrderedList(tt.line)
			if result != tt.expected {
				t.Errorf("isOrderedList(%q) = %v, want %v", tt.line, result, tt.expected)
			}
		})
	}
}
