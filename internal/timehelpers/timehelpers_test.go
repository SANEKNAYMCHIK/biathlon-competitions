package timehelpers

import "testing"

func TestMilliToTime(t *testing.T) {
	tests := []struct {
		name     string
		a        uint32
		expected string
	}{
		{"simple test", 47030, "00:00:47.030"},
		{"medium test", 72345678, "20:05:45.678"},
		{"hard test", 86399000, "23:59:59.000"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MilliToTime(tt.a); got != tt.expected {
				t.Errorf("MilliToTime(%v) = %v; want %s", tt.a, got, tt.expected)
			}
		})
	}
}
