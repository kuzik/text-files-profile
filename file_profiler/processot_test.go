package file_profiler_test

import (
	"github.com/kuzik/text-files-profile/file_profiler"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestProcessStats(t *testing.T) {

	fileStats := []file_profiler.FileStat{{1, 0, 0, 0, 5}, {3, 5, 6, 2, 4}}
	stats := make(chan file_profiler.FileStat)

	want := []file_profiler.RowStat{
		{4, 2},
		{5, 2},
		{6, 2},
		{2, 2},
		{9, 2},
	}

	go func(stats chan<- file_profiler.FileStat) {
		for _, fileStat := range fileStats {
			stats <- fileStat
		}

		close(stats)
	}(stats)

	processor := file_profiler.Processor{}
	got := processor.ProcessStats(stats)
	assert.Equal(t, want, got)
}
