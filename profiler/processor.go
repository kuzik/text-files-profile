package profiler

type Processor struct {
}

func (p Processor) ProcessStats(stats <-chan FileStat) []RowStat {
	rowStats := make([]RowStat, 1)
	for fileStat := range stats {
		for rowNumber, rowLength := range fileStat {
			if len(rowStats) <= rowNumber || len(rowStats) == 0 {
				rowStats = append(rowStats, RowStat{
					Sum:   rowLength,
					Count: 1,
				})
			} else {
				if rowLength != 0 {
					rowStats[rowNumber].Count++
				}
				rowStats[rowNumber].Sum += rowLength
			}
		}
	}

	return rowStats
}
