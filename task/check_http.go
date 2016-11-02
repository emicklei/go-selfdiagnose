package task

// Copyright 2013 Ernest Micklei. All rights reserved.
// Use of this source code is governed by a license
// that can be found in the LICENSE file.

import (
	"bytes"
	"fmt"
	"io"
	"net/http"

	. "github.com/emicklei/go-selfdiagnose"
)

// CheckHttp send a http.Request and check the status code. 200 OK = Passed
type CheckHttp struct {
	BasicTask
	Request      *http.Request
	ShowResponse bool
	HTTPClient   *http.Client
}

// Run sends the request and updates the result.
func (c CheckHttp) Run(ctx *Context, result *Result) {
	client := c.HTTPClient
	if client == nil {
		client = http.DefaultClient
	}
	resp, err := client.Do(c.Request)
	if err != nil {
		result.Passed = false
		result.Reason = fmt.Sprintf("%s %s => %s", c.Request.Method, c.Request.URL.String(), err.Error())
		return
	}
	defer resp.Body.Close()
	result.Passed = resp.StatusCode == http.StatusOK
	summary := fmt.Sprintf("%s %s => %s", c.Request.Method, c.Request.URL.String(), resp.Status)
	if c.ShowResponse {
		var buf bytes.Buffer
		buf.WriteString(summary)
		buf.WriteString("\n\n")
		io.Copy(&buf, resp.Body)
		summary = buf.String()
	}
	result.Reason = summary
}
