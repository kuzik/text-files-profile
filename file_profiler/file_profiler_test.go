package file_profiler_test

import (
	"github.com/kuzik/text-files-profile/file_profiler"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewProfiler(t *testing.T) {

	profiler := file_profiler.NewProfiler(
		file_profiler.Collector{},
		file_profiler.Processor{},
	)

	assert.IsType(t, &file_profiler.Profiler{}, profiler)
}

func TestRowStat_Len(t *testing.T) {
	tests := []struct {
		name        string
		sum         int
		count       int
		expectedLen int
	}{
		{"zero count", 0, 0, 0},
		{"single symbol", 3, 3, 1},
		{"long row", 3, 1, 3},
		{"round test", 3, 2, 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rowStat := file_profiler.RowStat{
				Sum:   tt.sum,
				Count: tt.count,
			}

			assert.Equal(t, tt.expectedLen, rowStat.Len())
		})
	}
}
