package db

import (
	"fmt"
	"net/url"
	"shoollove/config"

	"github.com/gofiber/fiber/v2/log"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var Gorm *gorm.DB

func Init() {

	database := config.Config.Database

	option := &gorm.Config{
		// 配置查询时表非复数形式
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		Logger: logger.Default.LogMode(logger.Error),
	}

	var dialector gorm.Dialector

	if database.Driver == "postgres" {
		dialector = postgres.Open(fmt.Sprintf(
			"user=%s password=%s host=%s port=%d dbname=%s TimeZone=%s",
			database.User,
			database.Passwd,
			database.Addr,
			database.Port,
			database.DBName,
			database.Zone,
		))
	} else {
		dialector = mysql.Open(fmt.Sprintf(
			"%s:%s@tcp(%s:%d)/%s?loc=%s",
			database.User,
			database.Passwd,
			database.Addr,
			database.Port,
			database.DBName,
			url.QueryEscape(database.Zone),
		))
	}

	// 连接数据库
	db, err := gorm.Open(dialector, option)

	if err != nil {
		log.Error("连接数据库：%v", err)
	}

	// Gorm = db.Debug()
	Gorm = db
}
