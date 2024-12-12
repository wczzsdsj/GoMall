package dal

import (
	"gomall/demo/demo_proto/biz/dal/mysql"
)

func Init() {
	// redis.Init()
	mysql.Init()
}
