package rst

import (
	"testing"

	"github.com/sebdah/goldie/v2"
)

func TestTitleString(t *testing.T) {
	tests := []struct {
		name  string
		title title
	}{
		{
			name: "sample tititle",
			title: title{
				text: "title",
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.title.String()
			g := goldie.New(t)
			g.Assert(t, tc.name, []byte(got))
		})
	}
}
