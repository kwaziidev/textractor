package textractor

import "testing"

func Test_matchAuthor(t *testing.T) {
	tests := []struct {
		text  string
		want  string
		want1 bool
	}{
		{
			text:  "责任编辑：张某",
			want:  "张某",
			want1: true,
		},
		{
			text:  "原创:张某",
			want:  "张某",
			want1: true,
		},
		{
			text:  "作者:张某",
			want:  "张某",
			want1: true,
		},
		{
			text:  "作者张某",
			want:  "",
			want1: false,
		},
	}
	for _, tt := range tests {
		got, got1 := matchAuthor(tt.text)
		if got != tt.want {
			t.Errorf("matchAuthor() got = %v, want %v, input %v", got, tt.want, tt.text)
		}
		if got1 != tt.want1 {
			t.Errorf("matchAuthor() got1 = %v, want %v, input %v", got1, tt.want1, tt.text)
		}
	}
}
