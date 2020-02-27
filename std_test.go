package textractor

import "testing"

func Test_std(t *testing.T) {
	tests := []struct {
		input []float64
		want  float64
	}{
		{
			input: []float64{13, 23, 12, 44, 55},
			want:  17.21162397916013,
		},
	}
	for _, tt := range tests {
		if got := std(tt.input); got != tt.want {
			t.Errorf("std() = %v, want %v, input %v", got, tt.want, tt.input)
		}
	}
}
