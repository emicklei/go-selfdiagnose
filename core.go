package selfdiagnose

// Copyright 2013 Ernest Micklei. All rights reserved.
// Use of this source code is governed by a license
// that can be found in the LICENSE file.

const VERSION = "go-selfdiagnose 1.0"

// Task describes a diagnostic task that can be run.
type Task interface {
	Run(ctx *Context, result *Result)
	Description() string
}

// Result captures the execution result of a Task.
type Result struct {
	Target Task
	Passed bool
	Reason string
}

// Context can be used to read/write variable during the execution of a selfdiagnose run.
type Context struct {
	Variables map[string]interface{}
}

func newContext() *Context {
	return &Context{map[string]interface{}{}}
}

// CommentHolder is used to embed a comment.
type CommentHolder struct {
	Comment string
}

func (h CommentHolder) Description() string {
	return h.Comment
}

// Reporter describes how to report task execution results.
type Reporter interface {
	Report(results []*Result)
}
