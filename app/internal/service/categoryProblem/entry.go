package categoryProblem

type Group struct {
}

func (g *Group) CategoryProblemService() *sCategoryProblem {
	return newCategoryProblemService()
}
