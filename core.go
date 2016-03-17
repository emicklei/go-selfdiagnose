package selfdiagnose

// Copyright 2015 Ernest Micklei. All rights reserved.
// Use of this source code is governed by a license
// that can be found in the LICENSE file.

import "time"

const VERSION = "go-selfdiagnose 1.2"

var since = time.Now()

// Task describes a diagnostic task that can be run.
type Task interface {
	Run(ctx *Context, result *Result)
	Comment() string
}

type HasTimeout interface {
	Timeout() time.Duration
}

// Result captures the execution result of a Task.
type Result struct {
	Target      Task
	Passed      bool
	Reason      string
	CompletedIn time.Duration
	Severity    Severity
}

// Context can be used to read/write variable during the execution of a selfdiagnose run.
type Context struct {
	Variables map[string]interface{}
}

type Severity string

type HasSeverity interface {
	Severity() Severity
}

const (
	SeverityNone     Severity = "none"
	SeverityWarning  Severity = "warning"
	SeverityCritical Severity = "critical"
)

type BasicTask struct {
	comment  string
	timeout  time.Duration
	severity Severity
}

func (t BasicTask) Comment() string {
	return t.comment
}

func (t *BasicTask) SetComment(text string) {
	t.comment = text
}

func (t BasicTask) Timeout() time.Duration {
	return t.timeout
}

func (t *BasicTask) SetTimeout(after time.Duration) {
	t.timeout = after
}

func (t *BasicTask) SetSeverity(s Severity) {
	t.severity = s
}

func (t BasicTask) Severity() Severity {
	if len(t.severity) == 0 {
		return SeverityCritical
	}
	return t.severity
}

// NewContext creates a new empty Context to run tasks.
func NewContext() *Context {
	return &Context{map[string]interface{}{}}
}

// Reporter describes how to report task execution results.
type Reporter interface {
	Report(results []*Result)
}
