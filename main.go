package main

import (
	"fmt"
	"github.com/kuzik/text-files-profile/file_profiler"
	"os"
	"strings"
)

func main() {
	dir := os.Args[1]
	if dir == "" {
		panic("Missed required parameter")
	}

	profile, err := file_profiler.Profile(dir)
	if err != nil {
		panic("Error during profiling process")
	}

	for rowNumber, row := range profile {
		if rowNumber == 0 {
			continue
		}
		fmt.Printf("%v: %s\n", rowNumber, strings.Repeat("*", row.Len()))
	}
}
