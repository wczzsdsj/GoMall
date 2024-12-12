package dal

import (
	"gomall/demo/demo_thrift/biz/dal/mysql"
	"gomall/demo/demo_thrift/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
