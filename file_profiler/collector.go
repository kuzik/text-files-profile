package file_profiler

import (
	"bufio"
	"io/ioutil"
	"os"
	"sync"
	"unicode/utf8"
)

type Collector struct {
}

func (c Collector) CollectStat(dir string) <-chan FileStat {

	stats := make(chan FileStat)

	go func(stats chan<- FileStat, dir string) {
		var wg sync.WaitGroup
		items, _ := ioutil.ReadDir(dir)
		for _, item := range items {
			wg.Add(1)
			if item.IsDir() {
				go c.collectDirStat(stats, &wg, dir+"/"+item.Name())
			} else {
				go c.collectFileStat(stats, &wg, dir+"/"+item.Name())
			}
		}

		wg.Wait()
		close(stats)
	}(stats, dir)

	return stats
}

func (c Collector) collectDirStat(stats chan<- FileStat, wg *sync.WaitGroup, dir string) {
	items, _ := ioutil.ReadDir(dir)
	for _, item := range items {
		wg.Add(1)
		if item.IsDir() {
			go c.collectDirStat(stats, wg, dir+"/"+item.Name())
		} else {
			go c.collectFileStat(stats, wg, dir+"/"+item.Name())
		}
	}
	defer wg.Done()
}

func (c Collector) collectFileStat(stats chan<- FileStat, wg *sync.WaitGroup, filePath string) {
	file, _ := os.Open(filePath)
	defer file.Close()

	fileStat := make(FileStat, 1)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fileStat = append(fileStat, utf8.RuneCountInString(scanner.Text()))
	}

	stats <- fileStat
	defer wg.Done()
}
