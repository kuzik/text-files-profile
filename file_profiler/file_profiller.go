package file_profiler

type FileStat []int
type RowStat struct {
	sum   int
	count int
}

func (row RowStat) Len() int {
	if row.count == 0 {
		return 0
	}
	return row.sum / row.count
}

func Profile(dir string) ([]RowStat, error) {

	return ProcessStats(CollectStat(dir)), nil
}
