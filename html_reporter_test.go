package selfdiagnose

import "testing"
import "os"

func TestReportInHtml(t *testing.T) {
	reg := &Registry{}
	{
		check := CheckDirectory{Path: "/TestReportInHtml"}
		check.SetComment("test critical")
		reg.Register(check)
	}
	{
		check := CheckDirectory{Path: "/TestReportInHtml"}
		check.SetSeverity(SeverityWarning)
		check.SetComment("test warning")
		reg.Register(check)
	}
	{
		check := CheckDirectory{Path: os.TempDir()}
		check.SetComment("test none")
		reg.Register(check)
	}
	{
		check := CheckDirectory{Path: os.TempDir()}
		check.SetComment("test odd/even")
		reg.Register(check)
	}
	f, _ := os.Create("TestReportInHtml.html")
	defer f.Close()
	rep := HtmlReporter{f}
	reg.Run(rep)
}
