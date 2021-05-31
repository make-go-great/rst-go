// https://docutils.sourceforge.io/docs/user/rst/quickstart.html

package rst

import "strings"

const (
	titleToken = '='
	listToken  = '-'
)

var listTokens = map[rune]struct{}{
	listToken: {},
	'+':       {},
	'*':       {},
}

// Node is single rst syntax representation
// Example: title, list, ...
type Node interface {
	String() string
}

type Title struct {
	text string
}

func NewTitle(text string) Title {
	return Title{
		text: text,
	}
}

func (t Title) String() string {
	text := strings.TrimSpace(t.text)

	var builder strings.Builder

	for i := 0; i < len(text); i++ {
		builder.WriteString(string(titleToken))
	}

	builder.WriteString("\n")
	builder.WriteString(text)
	builder.WriteString("\n")

	for i := 0; i < len(text); i++ {
		builder.WriteString(string(titleToken))
	}

	return builder.String()
}

type Section struct {
	text string
}

func NewSection(text string) Node {
	return nil
}

type ListItem struct {
	text string
}

func NewListItem(text string) Node {
	return ListItem{
		text: text,
	}
}

func (i ListItem) String() string {
	text := strings.TrimSpace(i.text)

	return string(listToken) + " " + text
}

func Equal(n1, n2 Node) bool {
	return n1.String() == n2.String()
}
