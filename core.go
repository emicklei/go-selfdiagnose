package selfdiagnose

type Task interface {
	Run(ctx *Context, result *Result)
	Comment() string
}

type Result struct {
	Target       *Task
	Passed       bool
	FailedReason string
}

type Context struct {
	Variables map[string]interface{}
}

func NewContext() *Context {
	return &Context{map[string]interface{}{}}
}

type CommentHolder struct {
	Comment string
}

type Reporter interface {
	Report(results []*Result)
}
