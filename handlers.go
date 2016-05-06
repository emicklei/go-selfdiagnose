package selfdiagnose

// Copyright 2015 Ernest Micklei. All rights reserved.
// Use of this source code is governed by a license
// that can be found in the LICENSE file.

import (
	"net/http"
	"strings"
)

func AddInternalHandlers() {
	http.HandleFunc("/internal/selfdiagnose.html", handleSelfdiagnose)
	http.HandleFunc("/internal/selfdiagnose.xml", handleSelfdiagnose)
	http.HandleFunc("/internal/selfdiagnose.json", handleSelfdiagnose)
}

func handleSelfdiagnose(w http.ResponseWriter, r *http.Request) {
	ctx := NewContext()
	// prepare for ReportHttpRequest
	ctx.Variables["http.request"] = r
	var reporter Reporter
	if strings.HasSuffix(r.URL.Path, ".json") || r.URL.Query().Get("format") == "json" {
		w.Header().Set("Content-Type", "application/json")
		reporter = JSONReporter{w}
	} else if strings.HasSuffix(r.URL.Path, ".xml") || r.URL.Query().Get("format") == "xml" {
		w.Header().Set("Content-Type", "application/xml")
		reporter = XMLReporter{w}
	} else {
		w.Header().Set("Content-Type", "text/html")
		reporter = HtmlReporter{w}
	}
	DefaultRegistry.RunWithContext(reporter, ctx)
}
