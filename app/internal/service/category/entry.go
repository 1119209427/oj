package category

type Group struct{}

func (g *Group) CategoryService() *sCategory {
	return newSCategory()
}
