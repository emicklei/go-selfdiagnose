package selfdiagnose

// Copyright 2013 Ernest Micklei. All rights reserved.
// Use of this source code is governed by a license
// that can be found in the LICENSE file.

// TODO: maybe a separate Registry object to avoid globallness.

var tasks = []Task{}

// Register adds a task to the global registry
func Register(t Task) {
	tasks = append(tasks, t)
}

// Run executes all registered task (in order) and reports using a Reporter.
func Run(r Reporter) {
	ctx := newContext()
	results := []*Result{}
	for _, each := range tasks {
		res := new(Result)
		res.Target = each
		each.Run(ctx, res)
		results = append(results, res)
	}
	r.Report(results)
}
