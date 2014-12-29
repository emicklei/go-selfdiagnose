package selfdiagnose

// Copyright 2013 Ernest Micklei. All rights reserved.
// Use of this source code is governed by a license
// that can be found in the LICENSE file.

import "time"

var DefaultRegistry = Registry{}

// Registry holds the collection or registered Tasks. It can run them all.
type Registry struct {
	tasks []Task
}

// Register adds a task to the collection.
func (r *Registry) Register(t Task) {
	r.tasks = append(r.tasks, t)
}

// Run executes all registered task (in order) and reports using a Reporter.
func (r Registry) Run(rep Reporter) {
	r.RunWithContext(rep, NewContext())
}

func (r Registry) RunWithContext(rep Reporter, ctx *Context) {
	results := []*Result{}
	for _, each := range r.tasks {
		res := new(Result)
		res.Target = each
		now := time.Now()
		each.Run(ctx, res)
		res.CompletedIn = time.Now().Sub(now)
		results = append(results, res)
	}
	rep.Report(results)
}

// Register adds a task to the default registry
func Register(t Task) {
	DefaultRegistry.Register(t)
}

// Run delegates to the DefaultRegistry
func Run(rep Reporter) {
	DefaultRegistry.Run(rep)
}
