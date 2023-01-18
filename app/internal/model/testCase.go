package model

type TestCase struct {
	Id        int    `db:"id"`
	ProblemId int    `db:"problem_id"`
	Input     string `db:"input"`
	Output    string `db:"output"`
}
