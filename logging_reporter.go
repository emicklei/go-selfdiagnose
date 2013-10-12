package selfdiagnose

import "log"

type LoggingReporter struct{}

func (l LoggingReporter) Report(results []*Result) {
	for _, each := range results {
		if each.Passed {
			log.Printf("[ok] %s", each.Target.Comment())
		}
	}
}
