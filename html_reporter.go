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
}

type resultTable struct {
	Rows        []resultRow
	PassedCount int
	FailedCount int
	CompletedIn time.Duration
	Version     string
}

// Report produces a HTML report including a summary
func (h HtmlReporter) Report(results []*Result) {
	rows := []resultRow{}
	passedCount := 0
	failedCount := 0
	for _, each := range results {
		row := resultRow{}
		row.Description = each.Reason
		row.Comment = each.Target.Description()
		row.Passed = each.Passed
		if each.Passed {
			row.Verdict = "passed"
			passedCount++
		} else {
			row.Verdict = "failed"
			failedCount++
		}
		rows = append(rows, row)
	}
	resultTable := resultTable{Rows: rows, PassedCount: passedCount, FailedCount: failedCount, Version: VERSION}
	htmlTemplate.Execute(h.Writer, resultTable)
}

var htmlTemplate = template.Must(template.New("Page").Parse(`
<html>
<body>
	<table>
		<tr>
			<th>Check</th>
			<th>Comment</th>
			<th>Description</th>
		</tr>
		{{range .Rows}}
		<tr>
			<td>{{.Verdict}}</td>	
			<td>{{.Comment}}</td>	
			<td>{{.Description}}</td>	
		</tr>		
		{{end}}
	</table>
	
	<table>
		<tr>			
			<td>Passed:{{.PassedCount}}</td>
			<td>Failed:{{.PassedCount}}</td>
			<td>{{.Version}}</td>
		</tr>
	</table>
</body>
</html>`))
