package profiler

import (
	"bufio"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
	"unicode/utf8"
)

type Collector struct {
	Extensions []string
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
	defer wg.Done()

	items, _ := ioutil.ReadDir(dir)
	for _, item := range items {
		wg.Add(1)
		if item.IsDir() {
			go c.collectDirStat(stats, wg, dir+"/"+item.Name())
		} else {
			go c.collectFileStat(stats, wg, dir+"/"+item.Name())
		}
	}
}

func (c Collector) collectFileStat(stats chan<- FileStat, wg *sync.WaitGroup, path string) {
	defer wg.Done()

	if c.Extensions[0] != "*" && !c.supported(filepath.Ext(path)) {
		return
	}

	file, _ := os.Open(path)
	defer file.Close()
	fileStat := make(FileStat, 1)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fileStat = append(fileStat, utf8.RuneCountInString(scanner.Text()))
	}

	stats <- fileStat
}

func (c Collector) supported(ext string) bool {
	for _, v := range c.Extensions {
		if v == ext {
			return true
		}
	}
	return false
}
