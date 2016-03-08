package selfdiagnose

// Copyright 2016 Ernest Micklei. All rights reserved.
// Use of this source code is governed by a license
// that can be found in the LICENSE file.

import (
	"encoding/json"
	"fmt"
	"io"
	"time"
)

// JSONReporter is to produce a JSON report and it written on an io.Writer.
type JSONReporter struct {
	Writer io.Writer
}

// Report produces a HTML report including a summary
func (j JSONReporter) Report(results []*Result) {
	report := jsonReport{
		SelfDiagnose: map[string]string{
			"version": VERSION,
		},
		Run:     time.Now(),
		Results: j.buildResults(results),
	}
	json.NewEncoder(j.Writer).Encode(report)
}

func (j JSONReporter) buildResults(results []*Result) (list []jsonResult) {
	for _, each := range results {
		list = append(list, jsonResult{
			Task:     fmt.Sprintf("%T", each.Target),
			Status:   j.status(each.Passed),
			Comment:  each.Target.Comment(),
			Message:  each.Reason,
			Duration: each.CompletedIn.String(),
			Severity: string(each.Severity),
		})
	}
	return
}

func (j JSONReporter) status(ok bool) string {
	if !ok {
		return "FAIL"
	}
	return "OK"
}

type jsonReport struct {
	SelfDiagnose map[string]string `json:"selfdiagnose"`
	Run          time.Time         `json:"run"`
	Results      []jsonResult      `json:"results"`
}

type jsonResult struct {
	Task     string `json:"task"`
	Status   string `json:"status"`
	Comment  string `json:"comment"`
	Message  string `json:"message"`
	Duration string `json:"duration"`
	Severity string `json:"severity"`
}
