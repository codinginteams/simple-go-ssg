package content

import (
	"fmt"
	"regexp"
	"strings"
)

func MarkdownToHtml(input string) string {
	input = parseHeadings(input)
	input = parseBold(input)
	input = parseItalic(input)
	input = parseLinks(input)
	input = parseUnorderedList(input)
	input = parseOrderedList(input)
	return input
}

func parseHeadings(input string) string {
	return regexp.MustCompile(`(?m)^# (.+)$`).ReplaceAllString(input, "<h1>$1</h1>")
}

func parseBold(input string) string {
	return regexp.MustCompile(`\*\*(.+?)\*\*`).ReplaceAllString(input, "<strong>$1</strong>")
}

func parseItalic(input string) string {
	return regexp.MustCompile(`_(.+?)_`).ReplaceAllString(input, "<em>$1</em>")
}

func parseLinks(input string) string {
	if strings.Contains(input, "[") && strings.Contains(input, "](") && strings.Contains(input, ")") {
		textStart := strings.Index(input, "[") + 1
		textEnd := strings.Index(input, "]")
		if textEnd < textStart {
			return input
		}
		text := input[textStart:textEnd]

		urlStart := strings.Index(input, "(") + 1
		urlEnd := strings.Index(input, ")")
		if urlEnd < urlStart {
			return input
		}
		url := input[urlStart:urlEnd]

		return fmt.Sprintf(`<a href="%s">%s</a>`, url, text)
	}

	return input
}

func parseUnorderedList(input string) string {
	items := regexp.MustCompile(`(?m)^- (.+)$`).FindAllStringSubmatch(input, -1)
	if len(items) == 0 {
		return input
	}
	var listItems []string
	for _, item := range items {
		listItems = append(listItems, "<li>"+item[1]+"</li>")
	}
	return "<ul>" + strings.Join(listItems, "") + "</ul>"
}

func parseOrderedList(input string) string {
	items := regexp.MustCompile(`(?m)^\d+\. (.+)$`).FindAllStringSubmatch(input, -1)
	if len(items) == 0 {
		return input
	}
	var listItems []string
	for _, item := range items {
		listItems = append(listItems, "<li>"+item[1]+"</li>")
	}
	return "<ol>" + strings.Join(listItems, "") + "</ol>"
}
