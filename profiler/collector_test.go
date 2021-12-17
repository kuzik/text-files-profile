package profiler_test

import (
	"github.com/kuzik/text-files-profile/profiler"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestCollectStat(t *testing.T) {
	expected := []profiler.FileStat{
		{0, 4, 3, 2, 1, 1, 4, 0, 6},
		{0, 1, 2, 3, 4},
		{0, 7, 0, 2, 0, 0, 0, 0, 1, 1, 2, 3},
	}

	collector := profiler.Collector{Extensions: []string{".txt"}}
	current, _ := os.Getwd()
	got := collector.CollectStat(current + "/test")

	for stat := range got {
		if !assert.Contains(t, expected, stat) {
			t.Fatal(stat)
		}
	}
}
