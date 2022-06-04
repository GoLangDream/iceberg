package database

import (
	"fmt"
	"github.com/gookit/config/v2"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DBConn *gorm.DB

func Init() {

	switch config.String("database.adapter") {
	case "mysql":
		dsn := fmt.Sprintf(
			"%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local",
			config.String("database.username"),
			config.String("database.password"),
			config.String("database.host"),
			config.String("database.port"),
			config.String("database.name"),
			config.String("database.encoding"),
		)
		DBConn, _ = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	case "postgres":
		dsn := fmt.Sprintf(
			"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
			config.String("database.host"),
			config.String("database.username"),
			config.String("database.password"),
			config.String("database.name"),
			config.String("database.port"),
		)
		DBConn, _ = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	case "sqlite":
		DBConn, _ = gorm.Open(sqlite.Open(config.String("database.file")), &gorm.Config{})
	}
}
