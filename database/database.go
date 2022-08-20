package database

import (
	"context"
	"fmt"
	"github.com/GoLangDream/iceberg/apm"
	"github.com/GoLangDream/iceberg/environment"
	"github.com/GoLangDream/iceberg/log"
	"github.com/gookit/config/v2"
	_ "github.com/newrelic/go-agent/v3/integrations/nrmysql"
	"github.com/newrelic/go-agent/v3/newrelic"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DBConn *gorm.DB = nil
var newRelicDB *gorm.DB = nil
var gormTransactionTracer *newrelic.Transaction

func StartTrace() {
	if newRelicDB == nil {
		return
	}
	gormTransactionTracer = apm.App.StartTransaction("GORM Operation")
	gormTransactionContext := newrelic.NewContext(context.Background(), gormTransactionTracer)
	DBConn = newRelicDB.WithContext(gormTransactionContext)
}

func EndTrace() {
	if newRelicDB == nil {
		return
	}
	gormTransactionTracer.End()
}

func Init() {
	var err error = nil

	switch config.String(keyWithEnv("database.%s.adapter")) {
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
		if apm.App == nil {
			DBConn, _ = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		} else {
			newRelicDB, _ = gorm.Open(mysql.New(mysql.Config{
				DriverName: "nrmysql",
				DSN:        dsn,
			}), &gorm.Config{})
		}

	case "postgres":
		dsn := fmt.Sprintf(
			"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
			config.String(keyWithEnv("database.%s.host")),
			config.String(keyWithEnv("database.%s.username")),
			config.String(keyWithEnv("database.%s.password")),
			config.String(keyWithEnv("database.%s.name")),
			config.String(keyWithEnv("database.%s.port")),
		)
		DBConn, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	case "sqlite":
		DBConn, err = gorm.Open(
			sqlite.Open(
				config.String(keyWithEnv("database.%s.file")),
			),
			&gorm.Config{},
		)
	default:
		log.Info("没有可以加载的数据库驱动")
	}
	if err != nil {
		log.Infof("数据库链接失败 %s", err)
	}
}

func keyWithEnv(format string) string {
	return fmt.Sprintf(format, environment.Name())
}
