package selfdiagnose

// Copyright 2013 Ernest Micklei. All rights reserved.
// Use of this source code is governed by a license
// that can be found in the LICENSE file.

import (
	"html/template"
	"io"
	"time"
)

// HtmlReporter is to produce a HTML report and it written on an io.Writer.
type HtmlReporter struct {
	Writer io.Writer
}

type resultRow struct {
	Comment     string
	Description string
	Passed      bool
	Verdict     string
	RowStyle    string
}

type resultTable struct {
	Rows        []resultRow
	PassedCount int
	FailedCount int
	CompletedIn time.Duration
	Version     string
}

func (r resultTable) TotalCount() int {
	return r.PassedCount + r.FailedCount
}

// Report produces a HTML report including a summary
func (h HtmlReporter) Report(results []*Result) {
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
			row.Verdict = "passed"
			passedCount++
			if i%2 == 0 {
				row.RowStyle = "even"
			} else {
				row.RowStyle = "odd"
			}

		} else {
			row.Verdict = "failed"
			failedCount++
		}
		rows = append(rows, row)
		completedIn += each.CompletedIn
	}
	resultTable := resultTable{Rows: rows, PassedCount: passedCount, FailedCount: failedCount, CompletedIn: completedIn, Version: VERSION}
	htmlTemplate.Execute(h.Writer, resultTable)
}

var htmlTemplate = template.Must(template.New("Page").Parse(`
<html>
<body>
	<style>
		body, table {
			font-family:verdana;
			font-size:small;
		}
		.odd { background-color:#F3F5F8 }
		.even { background-color:#DCE2EB }
		.table {
			padding: 4px;
		}
		.passed { color: #000; }
		.failed { color: #0000ff; }
		.error { color: #ff0000; }	
	</style>
	<table>
		<tr class="odd">
			<th>Check</th>
			<th>Comment</th>
			<th>Description</th>
		</tr>
		{{range .Rows}}
		<tr class="{{.RowStyle}}">
			<td>{{.Verdict}}</td>	
			<td>{{.Comment}}</td>	
			<td>{{.Description}}</td>	
		</tr>		
		{{end}}
	</table>
	
	<table>
		<tr>			
			<td>Checks:{{.TotalCount}}</td>
			<td>Passed:{{.PassedCount}}</td>
			<td>Failures:{{.FailedCount}}</td>
			<td>Time:{{.CompletedIn}}</td>
			<td>Version:{{.Version}}</td>
		</tr>
	</table>
</body>
</html>`))
