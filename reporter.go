package selfdiagnose

import "time"

func buildResultTable(results []*Result) resultTable {
	rows := []resultRow{}
	passedCount := 0
	failedCount := 0
	completedIn := time.Duration(0)
	for i, each := range results {
		row := resultRow{}
		row.Description = each.Reason
		row.Comment = each.Target.Comment()
		row.Passed = each.Passed
		if each.Passed {
			row.DescriptionStyle = "passed"
			passedCount++
			if i%2 == 0 {
				row.RowStyle = "even"
			} else {
				row.RowStyle = "odd"
			}

		} else {
			row.DescriptionStyle = "failed"
			failedCount++
		}
		rows = append(rows, row)
		completedIn += each.CompletedIn
	}

	resultTable := resultTable{
		Rows:        rows,
		PassedCount: passedCount,
		FailedCount: failedCount,
		CompletedIn: completedIn,
		Version:     VERSION,
		ReportDate:  time.Now(),
	}
	return resultTable

}
