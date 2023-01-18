package model

import "time"

type Problem struct {
	Id         int       `db:"id"`
	Identity   string    `db:"identity"`
	MaxRuntime int       `db:"max_runtime"` //最大运行时间
	MaxMem     int       `db:"max_mem"`     //最大运行内存
	Title      string    `db:"title"`
	Content    string    `db:"content"`
	TextCase   string    `db:"text_case"`
	CreatedAt  time.Time `db:"created_at"`
}
