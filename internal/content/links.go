package content

import "strings"

// TODO: Refactor this file later to reduce the amount of code. It was not straightforward to refactor it at the moment.
func parseLinks(input string) string {
	var result strings.Builder
	start := 0
	length := len(input)

	for start < length {
		linkText := parseLinkText(input, start)
		linkURL := parseLinkURL(input, start)

		openBracket := strings.Index(input[start:], "[")
		if openBracket == -1 {
			result.WriteString(input[start:])
			break
		}
		openBracket += start

		closeParenthesis := strings.Index(input[start:], ")")
		if closeParenthesis == -1 {
			result.WriteString(input[start:])
			break
		}
		closeParenthesis += start

		result.WriteString(input[start:openBracket])
		result.WriteString(buildLink(linkText, linkURL))

		start = closeParenthesis + 1
	}

	return result.String()
}

func parseLinkText(input string, start int) string {
	openBracket := strings.Index(input[start:], "[")
	if openBracket == -1 {
		return ""
	}
	openBracket += start

	closeBracket := strings.Index(input[openBracket:], "]")
	if closeBracket == -1 {
		return ""
	}
	closeBracket += openBracket

	return input[openBracket+1 : closeBracket]
}

func parseLinkURL(input string, start int) string {
	closeBracket := strings.Index(input[start:], "]")
	if closeBracket == -1 {
		return ""
	}
	closeBracket += start

	openParenthesis := strings.Index(input[closeBracket:], "(")
	if openParenthesis == -1 {
		return ""
	}
	openParenthesis += closeBracket

	closeParenthesis := strings.Index(input[openParenthesis:], ")")
	if closeParenthesis == -1 {
		return ""
	}
	closeParenthesis += openParenthesis

	return input[openParenthesis+1 : closeParenthesis]
}

func buildLink(linkText, linkURL string) string {
	return `<a href="` + linkURL + `">` + linkText + `</a>`
}
