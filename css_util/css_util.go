package css_util

import (
	"regexp"
	"strings"
)

func Format(input string) string {
	// Remove leading and trailing whitespace
	input = strings.TrimSpace(input)

	// Normalize line endings
	input = strings.ReplaceAll(input, "\r\n", "\n")

	lines := strings.Split(input, "\n")
	var result []string
	indentLevel := 0
	lastWasCloseBrace := false

	for i := 0; i < len(lines); i++ {
		line := strings.TrimSpace(lines[i])

		// Skip empty lines
		if line == "" {
			continue
		}

		// Check if this is a closing brace
		if line == "}" {
			indentLevel--
			if indentLevel < 0 {
				indentLevel = 0
			}
			result = append(result, strings.Repeat("    ", indentLevel)+"}")
			lastWasCloseBrace = true
			continue
		}

		// Add blank line after closing brace at root level only
		if lastWasCloseBrace && indentLevel == 0 {
			result = append(result, "")
		}
		lastWasCloseBrace = false

		// Check if line opens a block
		opensBlock := strings.HasSuffix(line, "{")

		if opensBlock {
			// Handle comma-separated selectors in keyframes
			if strings.Contains(line, ",") && !strings.HasPrefix(line, "@") {
				// Split by comma and format each selector
				parts := strings.Split(line[:len(line)-1], ",")
				for j, part := range parts {
					part = strings.TrimSpace(part)
					if j < len(parts)-1 {
						result = append(result, strings.Repeat("    ", indentLevel)+part+",")
					} else {
						result = append(result, strings.Repeat("    ", indentLevel)+part+" {")
					}
				}
			} else {
				result = append(result, strings.Repeat("    ", indentLevel)+line)
			}
			indentLevel++
			continue
		}

		// Normalize property declarations (remove extra spaces around colons)
		if strings.Contains(line, ":") && !strings.HasPrefix(line, "@") {
			// This is a property declaration
			parts := strings.SplitN(line, ":", 2)
			if len(parts) == 2 {
				key := strings.TrimSpace(parts[0])
				value := strings.TrimSpace(parts[1])
				// Remove space before semicolon if present
				value = regexp.MustCompile(`\s+;`).ReplaceAllString(value, ";")
				line = key + ": " + value
			}
		}

		// Add the line with proper indentation
		result = append(result, strings.Repeat("    ", indentLevel)+line)
	}

	// Join all lines and ensure single trailing newline
	output := strings.Join(result, "\n")
	return output + "\n"
}
