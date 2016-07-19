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
