package submit

type Group struct{}

func (g *Group) Submit() *sSubmit {
	return NewSubmitService()
}
