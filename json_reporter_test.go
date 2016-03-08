package selfdiagnose

import "testing"
import "os"

func TestReportInJSON(t *testing.T) {
	reg := &Registry{}
	{
		check := CheckDirectory{Path: "/TestReportInJSON"}
		check.SetComment("test critical")
		reg.Register(check)
	}
	{
		check := CheckDirectory{Path: "/TestReportInJSON"}
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
	f, _ := os.Create("TestReportInJSON.json")
	defer f.Close()
	rep := JSONReporter{f}
	reg.Run(rep)
}
