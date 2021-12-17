package profiler_test

import (
	"github.com/kuzik/text-files-profile/profiler"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestProcessStats(t *testing.T) {

	fileStats := []profiler.FileStat{{1, 0, 0, 0, 5}, {3, 5, 6, 2, 4}}
	stats := make(chan profiler.FileStat)

	want := []profiler.RowStat{
		{4, 2},
		{5, 2},
		{6, 2},
		{2, 2},
		{9, 2},
	}

	go func(stats chan<- profiler.FileStat) {
		for _, fileStat := range fileStats {
			stats <- fileStat
		}

		close(stats)
	}(stats)

	processor := profiler.Processor{}
	got := processor.ProcessStats(stats)
	assert.Equal(t, want, got)
}
