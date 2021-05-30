// https://docutils.sourceforge.io/docs/user/rst/quickstart.html

package rst

const (
	titleToken       = '='
	defaultListToken = '-'
)

var alternativeListTokens = map[rune]struct{}{
	'+': {},
	'*': {},
}
