package model

import (
	"testing"
)

func TestColumnWidths(t *testing.T) {
	tests := []struct {
		name       string
		ratios     []int
		totalWidth int
		want       []int
	}{
		{
			name:       "[1,1] splits evenly",
			ratios:     []int{1, 1},
			totalWidth: 82, // 82 - 2 gap = 80 available, 40 each
			want:       []int{40, 40},
		},
		{
			name:       "[3,2] proportional",
			ratios:     []int{3, 2},
			totalWidth: 52, // 52 - 2 gap = 50 available; 3/5*50=30, remainder=20
			want:       []int{30, 20},
		},
		{
			name:       "empty ratios returns nil",
			ratios:     []int{},
			totalWidth: 80,
			want:       nil,
		},
		{
			name:       "very narrow width gives minimum 1 per column",
			ratios:     []int{1, 1, 1},
			totalWidth: 2, // 2 - 4 gaps = -2 available, less than 3 columns
			want:       []int{1, 1, 1},
		},
		{
			name:       "three columns",
			ratios:     []int{1, 2, 1},
			totalWidth: 44, // 44 - 4 gaps = 40 available; 1/4*40=10, 2/4*40=20, remainder=10
			want:       []int{10, 20, 10},
		},
		{
			name:       "single column gets full width",
			ratios:     []int{1},
			totalWidth: 80, // 0 gaps, 80 available
			want:       []int{80},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			layout := ColumnLayout{Ratios: tt.ratios}
			got := layout.ColumnWidths(tt.totalWidth)

			if tt.want == nil {
				if got != nil {
					t.Fatalf("ColumnWidths() = %v, want nil", got)
				}
				return
			}

			if len(got) != len(tt.want) {
				t.Fatalf("ColumnWidths() returned %d widths, want %d", len(got), len(tt.want))
			}

			for i := range got {
				if got[i] != tt.want[i] {
					t.Errorf("ColumnWidths()[%d] = %d, want %d", i, got[i], tt.want[i])
				}
			}
		})
	}
}
