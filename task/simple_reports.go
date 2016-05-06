package task

import (
	"fmt"
	"runtime"

	"github.com/emicklei/go-selfdiagnose"
)

func ReportCPU() selfdiagnose.ReportMessage {
	cpu := selfdiagnose.ReportMessage{
		Message: fmt.Sprintf("%d CPU available. %d goroutines active", runtime.NumCPU(), runtime.NumGoroutine()),
	}
	cpu.SetComment("Num CPU")
	return cpu
}

type ReportBuildAndDate struct {
	Build string
	Date  string
}

func (r ReportBuildAndDate) Run(ctx *selfdiagnose.Context, result *selfdiagnose.Result) {
	result.Passed = true
	result.Reason = fmt.Sprintf("version:%s date:%s", r.Build, r.Date)
}

func (r ReportBuildAndDate) Comment() string {
	return "build information"
}
