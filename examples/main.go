package main

import (
	"log"
	"net/http"

	"github.com/emicklei/go-selfdiagnose"
	"github.com/emicklei/go-selfdiagnose/task"
)

func main() {
	// add http handlers for /internal/selfdiagnose.(html|json|xml)
	selfdiagnose.AddInternalHandlers()

	// add tasks
	selfdiagnose.Register(task.ReportHttpRequest{})
	selfdiagnose.Register(task.ReportHostname{})
	// add custom task
	selfdiagnose.Register(CheckGithub{})
	selfdiagnose.Register(task.ReportDiskusage{Path: "/usr/local/bin"}) // works on linux, darwin, freebsd, openbsd, windows

	// fire up
	log.Println("open http://localhost:9292/internal/selfdiagnose.html")
	log.Fatal(http.ListenAndServe(":9292", nil))
}

type CheckGithub struct{}

func (c CheckGithub) Comment() string { return "check availability of github.com" }

func (c CheckGithub) Run(ctx *selfdiagnose.Context, result *selfdiagnose.Result) {
	// NOTE: wrong url on purpose to see failure
	_, err := http.Get("https://githup.com")
	result.Passed = err == nil
	if err != nil {
		result.Reason = err.Error()
	}
}
