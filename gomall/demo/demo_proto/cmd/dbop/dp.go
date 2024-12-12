package main

import (
	"fmt"
	"gomall/demo/demo_proto/biz/dal"
	"gomall/demo/demo_proto/biz/dal/mysql"
	"gomall/demo/demo_proto/biz/model"
)

func main() {
	dal.Init()
	mysql.DB.Create(&model.User{Email: "test@example", Password: "111"})
	mysql.DB.Model(&model.User{}).Where("email=?", "test@example").Update("password", "222")
	var row model.User
	mysql.DB.First(&model.User{}).Where("email=?", "test@example").First(&row)
	fmt.Printf("row:%+v\n", row)
	mysql.DB.Where("email=?", "test@example").Delete(&model.User{})
	// mysql.DB.Unscoped().Where("email = ?", "test@example").Delete(&model.User{})
}
