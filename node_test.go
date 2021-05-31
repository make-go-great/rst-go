package rst

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTitleString(t *testing.T) {
	tests := []struct {
		name  string
		title title
		want  string
	}{
		{
			name: "sample title",
			title: title{
				text: "title",
			},
			want: "=====\ntitle\n=====",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.title.String()
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestListItemString(t *testing.T) {
	tests := []struct {
		name     string
		listItem listItem
		want     string
	}{
		{
			name: "sample list item",
			listItem: listItem{
				text: "item A",
			},
			want: "- item A",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.listItem.String()
			assert.Equal(t, tc.want, got)
		})
	}
}
