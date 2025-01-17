package text

import (
	"strings"
)

// BuildHistory takes commit messages and builds a complete list
func BuildHistory(messages []Commit) string {
	builder := strings.Builder{}
	builder.WriteString("\n")

	for i := 0; i < len(messages); i++ {
		builder.WriteString("- ")
		builder.WriteString(messages[i].Category)
		builder.WriteString("   ")
		builder.WriteString(messages[i].Heading)
		builder.WriteString("\n")
	}

	return builder.String()
}
