package task

import (
	"net/http"
	"testing"

	. "github.com/emicklei/go-selfdiagnose"
)

func TestCheckHTTP(t *testing.T) {
	get, err := http.NewRequest("GET", "http://ernestmicklei.com", nil)
	if err != nil {
		t.Fatal(err)
	}
	check := NewCheckHTTP(get)
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

func ExampleCheckHTTP() {
	get, _ := http.NewRequest("GET", "http://ernestmicklei.com", nil)
	check := NewCheckHTTP(get)
	check.SetComment("blog access")
	Register(check)
}
