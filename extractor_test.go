package textractor

import (
	"strings"
	"testing"

	"github.com/PuerkitoBio/goquery"
	"github.com/stretchr/testify/assert"
)

func Test_emptyNode(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  bool
	}{
		{
			input: `<div></div>`,
			want:  false,
		},
		{
			input: `<h1></h1>`,
			want:  true,
		},
		{
			input: `<h1>test</h1>`,
			want:  false,
		},
		{
			input: `<section></section>`,
			want:  true,
		},
		{
			input: `<span></span>`,
			want:  true,
		},
		{
			input: `<span>        </span>`,
			want:  true,
		},
		{
			input: `<span>   
			     </span>`,
			want: true,
		},
		{
			input: `<span><b></b></span>`,
			want:  false,
		},
	}
	for _, tt := range tests {
		dom, err := goquery.NewDocumentFromReader(strings.NewReader(tt.input))
		assert.Nil(t, err, "")
		if got := canBeRemove(dom.Find("body").Children()); got != tt.want {
			t.Errorf("emptyNode() = %v, want %v, input \"%v\"", got, tt.want, tt.input)
		}
	}
}
