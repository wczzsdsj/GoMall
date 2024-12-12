package dal

import (
	"golang/demo/demo_proto/biz/dal/mysql"
	"golang/demo/demo_proto/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
