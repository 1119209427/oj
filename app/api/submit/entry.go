package submit

type Group struct {
}

func (g *Group) Submit() *Api {
	return &insSubmit
}
