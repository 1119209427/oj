package model

import "time"

type Category struct {
	Id        int       `db:"id"`
	Name      string    `db:"name"` //分类的名称
	ParentId  int       `db:"parent_id"`
	CreatedAt time.Time `db:"create_time"`
}
