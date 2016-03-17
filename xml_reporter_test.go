package selfdiagnose

import "testing"
import "os"

func TestReportInXML(t *testing.T) {
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
	f, _ := os.Create("TestReportInXML.xml")
	defer f.Close()
	rep := XMLReporter{f}
	reg.Run(rep)
}
