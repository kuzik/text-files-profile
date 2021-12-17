package main

import (
	"flag"
	"os"
	"strings"

	"github.com/kuzik/text-files-profile/profiler"
)

func main() {
	currentDir, _ := os.Getwd()
	fileExtensions := flag.String("extension", "*", "List of supported file extensions (coma separated)")
	dir := flag.String("dir", currentDir, "Base working dir (current dir by default)")
	flag.Parse()

	prof := profiler.NewProfiler(
		&profiler.Collector{
			Extensions: strings.Split(*fileExtensions, ","),
		},
		&profiler.Processor{},
	)

	profile, err := prof.Profile(*dir)
	if err != nil {
		panic("Error during profiling process")
	}

	prof.PrintProfile(profile)
}
