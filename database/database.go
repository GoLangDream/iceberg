package database

import (
	"fmt"
	"github.com/gookit/config/v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DBConn *gorm.DB

func InitDatabase() {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local",
		config.String("database.username"),
		config.String("database.password"),
		config.String("database.host"),
		config.String("database.port"),
		config.String("database.name"),
		config.String("database.encoding"),
	)
	switch config.String("database.adapter") {
	case "mysql":
		DBConn, _ = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	}
}
