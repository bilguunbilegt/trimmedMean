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
			name:        "Trim 5% from both ends",
			numbers:     []float64{1, 2, 3, 4, 5, 6, 7},
			trim:        []int{5, 5},
			expected:    4.0,
			expectError: false,
		},
		{
			name:        "Trim 10% from both ends",
			numbers:     []float64{1, 2, 3, 4, 5, 6, 7},
			trim:        []int{10, 10},
			expected:    4.0,
			expectError: false,
		},
		{
			name:        "Trim 20% from both ends",
			numbers:     []float64{1, 2, 3, 4, 5, 6, 7},
			trim:        []int{20, 20},
			expected:    4.0,
			expectError: false,
		},
		{
			name:        "Trim 40% from both ends",
			numbers:     []float64{1, 2, 3, 4, 5, 6, 7},
			trim:        []int{40, 40},
			expected:    4.0,
			expectError: false,
		},
		{
			name:        "Trim more than available elements, expect error",
			numbers:     []float64{1, 2, 3, 4, 5},
			trim:        []int{50, 50},
			expected:    0.0,
			expectError: true,
		},
		{
			name:        "Trim with invalid percentage, expect error",
			numbers:     []float64{1, 2, 3, 4, 5},
			trim:        []int{-10}, // Negative percentage
			expected:    0.0,
			expectError: true,
		},
		{
			name:        "Trim with invalid percentage, expect error",
			numbers:     []float64{1, 2, 3, 4, 5},
			trim:        []int{110}, // Percentage over 100
			expected:    0.0,
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := TrimmedMean(tt.numbers, tt.trim...)

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
