package category

type Group struct {
}

func (g *Group) Category() *CateApi {
	return &insCate
}
