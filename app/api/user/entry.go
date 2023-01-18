package user

type Group struct {
}

func (g *Group) User() *SignApi {
	return &insSign
}

func (g *Group) Admin() *AdminApi {
	return &insAdmin
}
