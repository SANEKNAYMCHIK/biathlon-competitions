package competition

import "testing"

func TestValsToTuple(t *testing.T) {
	tests := []struct {
		name     string
		a        []string
		b        []string
		expected string
	}{
		{"simple test", []string{"11:11:11.000"}, []string{"1.000"}, "{11:11:11.000, 1.000}"},
		{"test with empty vals", []string{"11:35:11.000", ""}, []string{"12.047", ""}, "{11:35:11.000, 12.047}, {,}"},
		{"big test with empty vals", []string{"11:11:11.000", "00:00:00.000", "", ""}, []string{"12.047", "56.030", "", ""}, "{11:11:11.000, 12.047}, {00:00:00.000, 56.030}, {,}, {,}"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := valsToTuples(tt.a, tt.b); got != tt.expected {
				t.Errorf("valsToTuples(%v, %v) = %v; want %s", tt.a, tt.b, got, tt.expected)
			}
		})
	}
}
