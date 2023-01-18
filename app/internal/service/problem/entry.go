package problem

type Group struct{}

func (g *Group) ProblemService() *sProblem {
	return newProblemService()
}
