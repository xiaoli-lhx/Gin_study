package main

import (
	"bubble/dao"
	"bubble/models"
	"bubble/routers"
)

func main() {

	//创建数据库
	//sql:CREATE DATABASE bubble;
	//连接数据库
	err := dao.InitMySQL()
	if err != nil {
		panic(err)
	}
	//关闭数据库 gorm2不需要手动关闭了

	//模型绑定
	dao.DB.AutoMigrate(&models.Todo{})
	r := routers.SetupRouter()
	_ = r.Run(":8080")
}
