package service

import (
	"oj/app/internal/service/category"
	"oj/app/internal/service/categoryProblem"
	"oj/app/internal/service/problem"
	"oj/app/internal/service/submit"
	"oj/app/internal/service/testCase"
	"oj/app/internal/service/user"
)

var insTestCase = testCase.Group{}

func TestCase() *testCase.Group {
	return &insTestCase
}

var insProblem = problem.Group{}

func Problem() *problem.Group {
	return &insProblem
}

var insCategoryProblem = categoryProblem.Group{}

func CategoryProblem() *categoryProblem.Group {
	return &insCategoryProblem
}

var insCategory = category.Group{}

func Category() *category.Group {
	return &insCategory
}

var insSubmit = submit.Group{}

func Submit() *submit.Group {
	return &insSubmit
}

var insUser = user.Group{}

func User() *user.Group {
	return &insUser
}
