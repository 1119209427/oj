package user

type Group struct{}

func (g *Group) User() *sUser {
	return NewUserService()
}
