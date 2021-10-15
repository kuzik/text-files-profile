package file_profiler

import (
	"fmt"
	"strings"
)

type FileStat []int
type RowStat struct {
	Sum   int
	Count int
}

func (row RowStat) Len() int {
	if row.Count == 0 {
		return 0
	}
	return row.Sum / row.Count
}

type Collectable interface {
	CollectStat(dir string) <-chan FileStat
}

type Processable interface {
	ProcessStats(stats <-chan FileStat) []RowStat
}

type Profiler struct {
	collector Collectable
	processor Processable
}

func (p Profiler) Profile(dir string) ([]RowStat, error) {

	return p.processor.ProcessStats(p.collector.CollectStat(dir)), nil
}

func (p Profiler) PrintProfile(rows []RowStat) {
	for rowNumber, row := range rows {
		if rowNumber == 0 {
			continue
		}
		fmt.Printf("%v: %s\n", rowNumber, strings.Repeat("*", row.Len()))
	}
}

func NewProfiler(collector Collectable, processor Processable) *Profiler {
	return &Profiler{
		collector: collector,
		processor: processor,
	}
}
