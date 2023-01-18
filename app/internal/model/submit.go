package model

import "time"

type Submit struct {
	Id        int       `db:"id"`
	Identity  string    `db:"identity"`
	ProblemId int       `db:"problem_id"`
	UserId    int       `db:"user_id"`
	Path      string    `db:"path"` //代码的存储路径
	Status    int       `db:"status"`
	CreatedAt time.Time `db:"created_at"'`
}
