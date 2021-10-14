package file_profiler

import (
	"bufio"
	"io/ioutil"
	"log"
	"os"
	"sync"
	"unicode/utf8"
)

func CollectStat(dir string) <-chan FileStat {

	stats := make(chan FileStat)

	go func(stats chan<- FileStat, dir string) {
		var wg sync.WaitGroup
		items, _ := ioutil.ReadDir(dir)
		for _, item := range items {
			wg.Add(1)
			if item.IsDir() {
				go collectDirStat(stats, &wg, dir+"/"+item.Name())
			} else {
				go collectFileStat(stats, &wg, dir+"/"+item.Name())
			}
		}

		wg.Wait()
		close(stats)
	}(stats, dir)

	return stats
}

func collectDirStat(stats chan<- FileStat, wg *sync.WaitGroup, dir string) {
	items, _ := ioutil.ReadDir(dir)
	for _, item := range items {
		wg.Add(1)
		if item.IsDir() {
			go collectDirStat(stats, wg, dir+"/"+item.Name())
		} else {
			go collectFileStat(stats, wg, dir+"/"+item.Name())
		}
	}
	defer wg.Done()
}

func collectFileStat(stats chan<- FileStat, wg *sync.WaitGroup, filePath string) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	fileStat := make(FileStat, 1)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fileStat = append(fileStat, utf8.RuneCountInString(scanner.Text()))
	}

	stats <- fileStat
	defer wg.Done()
}
