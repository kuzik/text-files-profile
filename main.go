package main

import (
	"github.com/kuzik/text-files-profile/file_profiler"
	"os"
)

func main() {
	dir := os.Args[1]
	if dir == "" {
		panic("Missed required parameter")
	}

	profiler := file_profiler.NewProfiler(
		&file_profiler.Collector{},
		&file_profiler.Processor{},
	)
	profile, err := profiler.Profile(dir)
	if err != nil {
		panic("Error during profiling process")
	}

	profiler.PrintProfile(profile)
}
