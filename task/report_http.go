package task

import (
	"bytes"
	"fmt"
	"net/http"
	"sort"

	"github.com/emicklei/go-selfdiagnose"
)

type ReportHttpRequest struct{}

func (r ReportHttpRequest) Run(ctx *selfdiagnose.Context, result *selfdiagnose.Result) {
	req, ok := ctx.Variables["http.request"]
	if !ok {
		result.Passed = false
		result.Reason = "missing variable 'http.request'"
		return
	}
	var buf bytes.Buffer
	// sort by key
	keys := []string{}
	headers := req.(*http.Request).Header
	for k := range headers {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		v := headers[k]
		buf.WriteString(fmt.Sprintf("%s = %s<br/>", k, v))
	}
	result.Passed = true
	result.Reason = buf.String()
}

func (r ReportHttpRequest) Comment() string { return "headers of this Http request" }
