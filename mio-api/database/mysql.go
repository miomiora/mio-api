package database

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"mio-api/config"
	"mio-api/model"
)

func InitMysql() {
	//连接MYSQL, 获得DB类型实例，用于后面的数据库读写操作。
	conn, err := gorm.Open(mysql.Open(config.Conf.Mysql.GetDsn()), &gorm.Config{
		SkipDefaultTransaction: true,
	})
	if err != nil {
		fmt.Println("[api init error]连接Mysql数据库失败, error=" + err.Error())
		return
	}
	// 连接成功
	fmt.Println("[Success] Mysql数据库连接成功！！！")
	err = conn.AutoMigrate(&model.User{}, &model.InterfaceInfo{}, &model.UserInterface{})
	if err != nil {
		fmt.Println("[database mysql] 创建表失败！")
	}
	db, err := conn.DB()
	if err != nil {
		fmt.Println("[database mysql] 获取sql实例失败！")
	}
	db.SetMaxIdleConns(config.Conf.Mysql.MaxConn)
	db.SetMaxOpenConns(config.Conf.Mysql.MaxOpen)

	DB = conn
}
