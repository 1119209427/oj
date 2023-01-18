package define

import "time"

const (
	DefaultPage       = "1"
	DefaultSize       = "20"
	DefaultRedisValue = -1 //redis中key对应的预设值，防脏读
	Admin             = 1
	CancelAdmin       = 0
	ParentId          = "0"
	ParentDefault     = 0
)

// OneDayOfHours 时间
var OneDayOfHours = 60 * 60 * 24
var OneMinute = 60 * 1
var OneMonth = 60 * 60 * 24 * 30
var OneYear = 365 * 60 * 60 * 24
var ExpireTime = time.Hour * 48 // 设置Redis数据热度消散时间。
