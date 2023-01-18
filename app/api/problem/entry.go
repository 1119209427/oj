package problem

type Group struct{}

func (g *Group) ProblemList() *ListApi {
	return &insList
}
func (g *Group) Problem() *ApiProblem {
	return &insProblem
}
func (g *Group) TestCase() *TestApi {
	return &insTestCase
}
