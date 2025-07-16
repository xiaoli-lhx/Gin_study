package main

//gorm_demo
import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func main() {
	//连接数据库
	dsn := "root:123456@tcp(127.0.0.1:3306)/db1?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	//迁移schema
	db.AutoMigrate(&Product{})
	//增删改查
	//p1 := &Product{
	//	Code:  "p1",
	//	Price: 1000,
	//}
	//db.Create(p1)
	//查询单条记录
	var p2 Product
	db.First(&p2, 1) //select * from products where id=1;
	//fmt.Println(p2)
	db.First(&p2, "code = ?", "p1") //select * from products where code="p1";
	//fmt.Printf("p2:%#v\n", p2)
	//更新
	//update - 将 product 的 price 更新为200
	//db.Debug().Model(&p2).Update("price", 200)可以看到执行的SQL语句
	db.Model(&p2).Update("price", 200)
	db.Model(&p2).Updates(map[string]interface{}{
		"Code":  "X42",
		"Price": 250,
	})
}
