package textractor

import "testing"

func Test_matchTime(t *testing.T) {
	tests := []struct {
		text  string
		want  string
		want1 bool
	}{
		{
			text:  "2020年02月27日 09:00",
			want:  "2020年02月27日 09:00",
			want1: true,
		},
		{
			text:  "2020年02月27日",
			want:  "2020年02月27日",
			want1: true,
		},
		{
			text:  "2020年02月27日 09:00:00",
			want:  "2020年02月27日 09:00:00",
			want1: true,
		},
		{
			text:  "2020-02-27 09:00:00",
			want:  "2020-02-27 09:00:00",
			want1: true,
		},
		{
			text:  "2020-02-27 09:00",
			want:  "2020-02-27 09:00",
			want1: true,
		},
		{
			text:  "2020-02-27",
			want:  "2020-02-27",
			want1: true,
		},
		{
			text:  "test 2020-02-27 test",
			want:  "2020-02-27",
			want1: true,
		},
		{
			text:  "2020-2-2 09:00",
			want:  "2020-2-2 09:00",
			want1: true,
		},
		{
			text:  "2-27 09:00",
			want:  "2-27 09:00",
			want1: true,
		},
		{
			text:  "2-27",
			want:  "2-27",
			want1: true,
		},
	}
	for _, tt := range tests {
		got, got1 := matchTime(tt.text)
		if got != tt.want {
			t.Errorf("matchTime() got = %v, want %v, input %v", got, tt.want, tt.text)
		}
		if got1 != tt.want1 {
			t.Errorf("matchTime() got1 = %v, want %v, input %v", got1, tt.want1, tt.text)
		}
	}
}
