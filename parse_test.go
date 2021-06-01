package rst

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseTitle(t *testing.T) {
	tests := []struct {
		name       string
		line       string
		wantResult Title
		wantOK     bool
	}{
		{
			name: "sample title",
			line: "=====\ntitle\n=====",
			wantResult: Title{
				text: "title",
			},
			wantOK: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			gotResult, gotOK := parseTitle(tc.line)
			assert.Equal(t, tc.wantOK, gotOK)
			if tc.wantOK {
				assert.Equal(t, tc.wantResult, gotResult)
			}
		})
	}
}

func TestParseSection(t *testing.T) {
	tests := []struct {
		name       string
		line       string
		wantResult Section
		wantOK     bool
	}{
		{
			name: "sample section",
			line: "section\n=======",
			wantResult: Section{
				text: "section",
			},
			wantOK: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			gotResult, gotOK := parseSection(tc.line)
			assert.Equal(t, tc.wantOK, gotOK)
			if tc.wantOK {
				assert.Equal(t, tc.wantResult, gotResult)
			}
		})
	}
}
