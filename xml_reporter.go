package selfdiagnose

import (
	"io"
	"log"
)

// HtmlReporter is to produce a HTML report and it written on an io.Writer.
type XMLReporter struct {
	Writer io.Writer
}

// Report produces an XML document
func (h XMLReporter) Report(results []*Result) {
	log.Println("in xml reporter")

}
