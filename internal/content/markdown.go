package content

import (
	"strings"
)

func MarkdownToHtml(input string) string {
	lines := strings.Split(input, "\n")
	var result []string
	var currentListType string
	var listBuffer []string

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		if isUnorderedList(line) {
			currentListType, listBuffer = handleFlatList(line, currentListType, "ul", listBuffer)
			continue
		}

		if isOrderedList(line) {
			currentListType, listBuffer = handleFlatList(line, currentListType, "ol", listBuffer)
			continue
		}

		if currentListType != "" {
			result = append(result, finalizeList(currentListType, listBuffer))
			currentListType = ""
			listBuffer = nil
		}

		if strings.HasPrefix(line, "#") {
			result = append(result, parseHeadings(line))
		} else {
			result = append(result, wrapWithParagraph(line))
		}
	}

	if currentListType != "" {
		result = append(result, finalizeList(currentListType, listBuffer))
	}

	return strings.Join(result, "")
}
