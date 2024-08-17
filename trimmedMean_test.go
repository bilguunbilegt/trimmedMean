package trimmedMean

import (
	"testing"
)

func TestTrimmedMean(t *testing.T) {
	tests := []struct {
		name        string
		numbers     []float64
		trim        []int
		expected    float64
		expectError bool
	}{
		{
			name:        "No trimming",
			numbers:     []float64{1, 2, 3, 4, 5, 6, 7},
			trim:        []int{0},
			expected:    4.0,
			expectError: false,
		},
		{
			name:        "Trim 1 from both ends",
			numbers:     []float64{1, 2, 3, 4, 5, 6, 7},
			trim:        []int{1, 1},
			expected:    4.0,
			expectError: false,
		},
		{
			name:        "Trim 2 from both ends",
			numbers:     []float64{1, 2, 3, 4, 5, 6, 7},
			trim:        []int{2, 2},
			expected:    4.0,
			expectError: false,
		},
		{
			name:        "Trim 3 from both ends",
			numbers:     []float64{1, 2, 3, 4, 5, 6, 7},
			trim:        []int{3, 3},
			expected:    4.0,
			expectError: false,
		},
		{
			name:        "Trim 4 from both ends, expect error",
			numbers:     []float64{1, 2, 3, 4, 5, 6, 7},
			trim:        []int{4, 4},
			expected:    0.0,
			expectError: true,
		},
		{
			name:        "Trim more elements than possible, expect error",
			numbers:     []float64{1, 2, 3, 4, 5},
			trim:        []int{3, 3},
			expected:    0.0,
			expectError: true,
		},
		{
			name:        "Invalid trim argument count, expect error",
			numbers:     []float64{1, 2, 3, 4, 5, 6, 7},
			trim:        []int{},
			expected:    0.0,
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			trimFloats := make([]float64, len(tt.trim))
			for i, v := range tt.trim {
				trimFloats[i] = float64(v)
			}
			result, err := TrimmedMean(tt.numbers, trimFloats...)
			if tt.expectError {
				if err == nil {
					t.Errorf("expected an error but got none")
				}
			} else {
				if err != nil {
					t.Errorf("expected no error but got one: %v", err)
				}
				if result != tt.expected {
					t.Errorf("expected %v, got %v", tt.expected, result)
				}
			}
		})
	}
}
