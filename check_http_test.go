package selfdiagnose

import (
	"net/http"
	"testing"
)

func TestCheckHttp(t *testing.T) {
	get, err := http.NewRequest("GET", "http://ernestmicklei.com", nil)
	if err != nil {
		t.Fatal(err)
	}
	check := CheckHttp{Request: get}
	check.SetComment("blog access")

	reg := &Registry{}
	reg.Register(check)

	rr := new(recordingReporter)
	reg.Run(rr)
	if len(rr.results) == 0 {
		t.Fatal("no results")
	}
	if got, want := rr.results[0].Passed, true; got != want {
		t.Errorf("got %v want %v", got, want)
	}
}

func ExampleCheckHttp() {
	get, _ := http.NewRequest("GET", "http://ernestmicklei.com", nil)
	check := CheckHttp{Request: get}
	check.SetComment("blog access")
	Register(check)
}
