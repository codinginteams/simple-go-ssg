// ./internal/content/headings.go
package content

import (
	"fmt"
	"regexp"
	"strings"
)

var headingRegex = regexp.MustCompile(`(?m)^#{1,6} (.+)$`)

func parseHeadings(input string) string {
	return headingRegex.ReplaceAllStringFunc(input, func(match string) string {
		level := len(strings.Split(match, " ")[0])
		return fmt.Sprintf("<h%d>%s</h%d>", level, match[level+1:], level)
	})
}
