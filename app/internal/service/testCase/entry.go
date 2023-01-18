package testCase

type Group struct{}

func (g *Group) TestCaseService() *sTestCase {
	return newTestCaseService()
}
