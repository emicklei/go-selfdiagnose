package task

// Copyright 2013,2016 Ernest Micklei. All rights reserved.
// Use of this source code is governed by a license
// that can be found in the LICENSE file.

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"sync"

	. "github.com/emicklei/go-selfdiagnose"
)

// CheckHTTP send a http.Request and check the status code. 200 OK = Passed
type CheckHTTP struct {
	BasicTask
	request      *http.Request
	mutex        *sync.Mutex
	ShowResponse bool
	HTTPClient   *http.Client
}

// NewCheckHTTP returns a CheckHTTP for threadsafe use of a Request.
func NewCheckHTTP(r *http.Request) *CheckHTTP {
	c := new(CheckHTTP)
	c.mutex = new(sync.Mutex)
	cp := *r
	if r.Body != nil {
		data, _ := ioutil.ReadAll(r.Body)
		cp.Body = ioutil.NopCloser(bytes.NewReader(data))
	}
	c.request = &cp
	return c
}

// Run sends the request and updates the result.
func (c *CheckHTTP) Run(ctx *Context, result *Result) {
	client := c.HTTPClient
	if client == nil {
		client = http.DefaultClient
	}
	// make copy of request
	request := *c.request
	if request.Body != nil {
		// safe copy body
		c.mutex.Lock()
		data, _ := ioutil.ReadAll(request.Body)
		request.Body = ioutil.NopCloser(bytes.NewReader(data))
		c.mutex.Unlock()
	}

	resp, err := client.Do(&request)
	if err != nil {
		result.Passed = false
		result.Reason = fmt.Sprintf("%s %s => %s", request.Method, request.URL.String(), err.Error())
		return
	}
	defer resp.Body.Close()
	result.Passed = resp.StatusCode == http.StatusOK
	summary := fmt.Sprintf("%s %s => %s", request.Method, request.URL.String(), resp.Status)
	if c.ShowResponse {
		var buf bytes.Buffer
		buf.WriteString(summary)
		buf.WriteString("\n\n")
		io.Copy(&buf, resp.Body)
		summary = buf.String()
	}
	result.Reason = summary
}
