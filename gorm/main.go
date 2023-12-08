package main

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

	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	// DSN (Data Source Name)
	dsn := "root:root1996@tcp(9.135.179.114:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	//db, err := gorm.Open(mysql.New(mysql.Config{
	//	DSN:                       dsn,   // DSN data source name
	//	DefaultStringSize:         256,   // string 类型字段的默认长度
	//	DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
	//	DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
	//	DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
	//	SkipInitializeWithVersion: false, // 根据当前 MySQL 版本自动配置
	//}), &gorm.Config{})
	println(db)
	println(err)

	//sqlDB, err := sql.Open("mysql", dsn)
	//gormDB, err := gorm.Open(mysql.New(mysql.Config{
	//	Conn: sqlDB,
	//}), &gorm.Config{})
	//println(gormDB)

	// 入门
	//db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	//if err != nil {
	//	panic("failed to connect database")
	//}
	//
	//// 迁移 schema
	//db.AutoMigrate(&Product{})
	//
	//// Create
	//db.Create(&Product{Code: "D42", Price: 100})
	//
	//// Read
	//var product Product
	//db.First(&product, 1)                 // 根据整型主键查找
	//db.First(&product, "code = ?", "D42") // 查找 code 字段值为 D42 的记录
	//
	//// Update - 将 product 的 price 更新为 200
	//db.Model(&product).Update("Price", 200)
	//// Update - 更新多个字段
	//db.Model(&product).Updates(Product{Price: 200, Code: "F42"}) // 仅更新非零值字段
	//db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})
	//
	//// Delete - 删除 product
	//db.Delete(&product, 1)
}
