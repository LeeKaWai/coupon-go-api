package db

import (
	"coupon-go-api/config"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB
var err error

func InitDB() {
	dbConfig := &config.GLOBAL_CONF.Mysql
	fmt.Println("database config:", dbConfig)

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		dbConfig.Username,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.Name,
	)

	fmt.Println("dsn:", dsn)

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		fmt.Println("连接数据库失败")
		panic(err)
	}

	fmt.Println("连接数据库成功")
}

func CloseDB() {
	if DB != nil {
		sqlDB, err := DB.DB()
		if err != nil {
			fmt.Println("获取底层数据库对象失败")
			panic(err)
		}
		sqlDB.Close()
	}
}
