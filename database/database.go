package database

import (
	"fmt"
	"github.com/GoLangDream/iceberg/environment"
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
			config.String(keyWithEnv("database.%s.username")),
			config.String(keyWithEnv("database.%s.password")),
			config.String(keyWithEnv("database.%s.host")),
			config.String(keyWithEnv("database.%s.port")),
			config.String(keyWithEnv("database.%s.name")),
			config.String(keyWithEnv("database.%s.encoding")),
		)
		DBConn, _ = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	case "postgres":
		dsn := fmt.Sprintf(
			"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
			config.String(keyWithEnv("database.%s.host")),
			config.String(keyWithEnv("database.%s.username")),
			config.String(keyWithEnv("database.%s.password")),
			config.String(keyWithEnv("database.%s.name")),
			config.String(keyWithEnv("database.%s.port")),
		)
		DBConn, _ = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	case "sqlite":
		DBConn, _ = gorm.Open(
			sqlite.Open(
				config.String(keyWithEnv("database.%s.file")),
			),
			&gorm.Config{},
		)
	}
}

func keyWithEnv(format string) string {
	return fmt.Sprintf(format, environment.Name())
}
