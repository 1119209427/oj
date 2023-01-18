package api

import (
	"oj/app/api/category"
	"oj/app/api/problem"
	"oj/app/api/submit"
	"oj/app/api/user"
)

var insCategory = category.Group{}

func Category() *category.Group {
	return &insCategory
}

var insProblem = problem.Group{}

func Problem() *problem.Group {
	return &insProblem
}

var insUser = user.Group{}

func User() *user.Group {
	return &insUser
}

var insSubmit = submit.Group{}

func Submit() *submit.Group {
	return &insSubmit
}
