package mysql

import (
	"gomall/demo/demo_proto/biz/model"
	"gomall/demo/demo_proto/conf"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func Init() {
	DB, err = gorm.Open(mysql.Open(conf.GetConf().MySQL.DSN),
		&gorm.Config{
			PrepareStmt:            true,
			SkipDefaultTransaction: true,
		},
	)
	// dsn := fmt.Sprintf(conf.GetConf().MySQL.DSN, os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASSWORD"), os.Getenv("MYSQL_HOST"), os.Getenv("MYSQL_DATABASE"))
	// DB, err = gorm.Open(mysql.Open(dsn),
	// 	&gorm.Config{
	// 		PrepareStmt:            true,
	// 		SkipDefaultTransaction: true,
	// 	},
	// )
	if err != nil {
		panic(err)
	}
	DB.AutoMigrate(&model.User{})
	// fmt.Printf("%#v", DB.Exec("select version()"))
}