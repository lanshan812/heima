package model

import (
	"fmt"
	"patent/util/log"
	"time"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/plugin/dbresolver"
)

var DB *gorm.DB

func InitDB() {
	addr := viper.GetString("DB.patent.addr")

	db, err := gorm.Open(mysql.Open(addr), &gorm.Config{DryRun: false}) //Logger: logger.Default.LogMode(logger.Silent)
	if err != nil {
		errInfo := fmt.Sprintf("%s", err)
		log.Logger.Error(errInfo)
	}

	//配置连接池
	db.Use(
		dbresolver.Register(dbresolver.Config{ /* xxx */ }).
			SetConnMaxIdleTime(time.Hour).SetConnMaxLifetime(24 * time.Hour).
			SetMaxIdleConns(100).
			SetMaxOpenConns(200),
	)

	DB = db
}
