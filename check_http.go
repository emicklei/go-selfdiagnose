package selfdiagnose

// Copyright 2013 Ernest Micklei. All rights reserved.
// Use of this source code is governed by a license
// that can be found in the LICENSE file.

import (
	"fmt"
	"net/http"
)

// CheckHttp send a http.Request and check the status code. 200 OK = Passed
type CheckHttp struct {
	CommentHolder
	Request *http.Request
}

// Run sends the request and updates the result.
func (c CheckHttp) Run(ctx *Context, result *Result) {
	client := new(http.Client)
	resp, err := client.Do(c.Request)
	defer resp.Body.Close()
	if err != nil {
		result.Passed = false
		result.Reason = err.Error()
		return
	}
	if resp.StatusCode != http.StatusOK {
		result.Passed = false
		result.Reason = resp.Status
		return
	}
	result.Passed = true
	result.Reason = fmt.Sprintf("%s %s => %s", c.Request.Method, c.Request.URL.String(), resp.Status)
}
