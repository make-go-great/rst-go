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
			name: "sample title",
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

func TestListItemString(t *testing.T) {
	tests := []struct {
		name     string
		listItem listItem
	}{
		{
			name: "sample list item",
			listItem: listItem{
				text: "item A",
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.listItem.String()
			g := goldie.New(t)
			g.Assert(t, tc.name, []byte(got))
		})
	}
}
