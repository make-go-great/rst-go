package rst

import "strings"

func GenerateText(nodes []Node) string {
	lines := make([]string, len(nodes))
	for i, node := range nodes {
		lines[i] = node.String()
	}

	result := strings.Join(lines, "\n\n")
	// Ensure there is newline at end of file
	result += "\n"

	return result
}
