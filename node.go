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

type title struct {
	text string
}

func (t title) String() string {
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
