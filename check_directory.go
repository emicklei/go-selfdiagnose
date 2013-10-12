package selfdiagnose

// CheckDirectory reports whether a directory (path) is readable and/or writable
type CheckDirectory struct {
	CommentHolder
	Path     string
	CanRead  bool
	CanWrite bool
}

func (c CheckDirectory) Run(ctx *Context, result *Result) {}
