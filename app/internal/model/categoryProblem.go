package model

type CategoryProblem struct {
	Id         int `db:"id"`
	CategoryId int `db:"category_id"`
	ProblemId  int `db:"problem_id"`
}
