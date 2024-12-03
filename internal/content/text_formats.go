// ./internal/content/text_formats.go
package content

import (
	"regexp"
)

var (
	boldItalicRegex = regexp.MustCompile(`\*\*_(.+?)_\*\*`)
	boldRegex       = regexp.MustCompile(`\*\*(.+?)\*\*`)
	italicRegex     = regexp.MustCompile(`_(.+?)_`)
)

func parseBoldItalic(input string) string {
	return boldItalicRegex.ReplaceAllString(input, `<strong><em>$1</em></strong>`)
}

func parseBold(input string) string {
	return boldRegex.ReplaceAllString(input, `<strong>$1</strong>`)
}

func parseItalic(input string) string {
	return italicRegex.ReplaceAllString(input, `<em>$1</em>`)
}

func wrapWithParagraph(line string) string {
	line = parseLinks(line)
	line = parseBoldItalic(line)
	line = parseBold(line)
	line = parseItalic(line)
	return "<p>" + line + "</p>"
}
