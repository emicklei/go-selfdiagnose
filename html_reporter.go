package selfdiagnose

// Copyright 2013 Ernest Micklei. All rights reserved.
// Use of this source code is governed by a license
// that can be found in the LICENSE file.

import (
	"io"
	"text/template"
	"time"
)

// HtmlReporter is to produce a HTML report and it written on an io.Writer.
type HtmlReporter struct {
	Writer io.Writer
}

type resultRow struct {
	Comment          string
	Description      string
	Passed           bool
	RowStyle         string
	DescriptionStyle string
}

type resultTable struct {
	Rows        []resultRow
	PassedCount int
	FailedCount int
	CompletedIn time.Duration
	Version     string
	ReportDate  time.Time
}

func (r resultTable) TotalCount() int {
	return r.PassedCount + r.FailedCount
}

// Report produces a HTML report including a summary
func (h HtmlReporter) Report(results []*Result) {
	resultTable := buildResultTable(results)
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
			<th>Comment</th>
			<th>Description</th>
		</tr>
		{{range .Rows}}
		<tr class="{{.RowStyle}}">
			<td>{{.Comment}}</td>	
			<td class="{{.DescriptionStyle}}">{{.Description}}</td>	
		</tr>		
		{{end}}
	</table>
	
	<h4>
		Checks: {{.TotalCount}} , Failures: {{.FailedCount}}, Time: {{.CompletedIn}} |
		{{.Version}}</td>
	</h4>
</body>
</html>`))
