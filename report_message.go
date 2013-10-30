package selfdiagnose

// ReportMessage simply does what it says. Can have comment too.
type ReportMessage struct {
	CommentHolder
	Message string
}

func (r ReportMessage) Run(ctx *Context, result *Result) {
	result.Passed = true
	result.Reason = r.Message
}
