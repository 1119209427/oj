package model

import "time"

type User struct {
	Id        int       `db:"id"`
	Admin     int       `db:"admin"`
	Identity  string    `db:"identity"`
	Name      string    `db:"name"`
	Password  string    `db:"password"`
	Salt      string    `db:"salt"`
	Phone     string    `db:"phone"`
	Mall      string    `db:"mall"`
	CreatedAt time.Time `db:"created_at"`
}
