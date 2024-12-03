// ./internal/content/lists.go
package content

import (
	"regexp"
	"strings"
)

var (
	unorderedListRegex = regexp.MustCompile(`^- `)
	orderedListRegex   = regexp.MustCompile(`^\d+\. `)
	listItemRegex      = regexp.MustCompile(`^[-\d\.]+\s`)
)

func isUnorderedList(line string) bool {
	return unorderedListRegex.MatchString(line)
}

func isOrderedList(line string) bool {
	return orderedListRegex.MatchString(line)
}

func handleFlatList(line, currentType, newType string, buffer []string) (string, []string) {
	item := parseListItemText(line)
	item = parseBold(parseItalic(parseLinks(item)))

	if currentType == newType {
		buffer = append(buffer, "<li>"+item+"</li>")
	} else {
		if currentType != "" {
			buffer = append(buffer, finalizeList(currentType, buffer))
			buffer = []string{"<li>" + item + "</li>"}
		} else {
			buffer = append(buffer, "<li>"+item+"</li>")
		}
		currentType = newType
	}
	return currentType, buffer
}

func finalizeList(listType string, listBuffer []string) string {
	return "<" + listType + ">" + strings.Join(listBuffer, "") + "</" + listType + ">"
}

func parseListItemText(line string) string {
	return strings.TrimSpace(listItemRegex.ReplaceAllString(line, ""))
}
