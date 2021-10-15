package file_profiler_test

import (
	"github.com/kuzik/text-files-profile/file_profiler"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestCollectStat(t *testing.T) {

	expected := []file_profiler.FileStat{
		{0, 4, 3, 2, 1, 1, 4, 0, 6},
		{0, 1, 2, 3, 4},
		{0, 7, 0, 2, 0, 0, 0, 0, 1, 1, 2, 3},
	}

	collector := file_profiler.Collector{}
	current, _ := os.Getwd()
	got := collector.CollectStat(current + "/test")

	for stat := range got {
		if !assert.Contains(t, expected, stat) {
			t.Fatal(stat)
		}
	}
}
