package file_profiler

func ProcessStats(stats <-chan FileStat) []RowStat {
	rowStats := make([]RowStat, 1)
	for fileStat := range stats {
		for rowNumber, rowLength := range fileStat {
			if len(rowStats) <= rowNumber || len(rowStats) == 0 {
				rowStats = append(rowStats, RowStat{
					sum:   rowLength,
					count: 1,
				})
			} else {
				rowStats[rowNumber].count++
				rowStats[rowNumber].sum += rowStats[rowNumber].sum + rowLength
			}
		}
	}

	return rowStats
}
