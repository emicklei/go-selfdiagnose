package selfdiagnose

var tasks = []*Task{}

func Register(t *Task) {
	tasks = append(tasks, t)
}

func Run(r *Reporter) {
	ctx := NewContext()
	results := []*Result{}
	for _, each := range tasks {
		res := new(Result)
		res.Target = each
		each.Run(ctx, res)
		results = append(results, res)
	}
}
