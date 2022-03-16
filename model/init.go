package model

import (
	"gin/config"
	"log"
	"time"

	"github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
)

// DB ...
var DB *xorm.Engine

// Database ...
func Database(config config.MySQLConfig) {
	d := mysql.NewConfig()
	d.User = config.User
	d.Passwd = config.Password
	d.Net = "tcp"
	d.Addr = config.Host + ":" + config.Port
	d.DBName = config.Name
	d.Params = map[string]string{
		"charset": config.Charset,
	}
	dataSource := d.FormatDSN()
	engine, err := xorm.NewEngine("mysql", dataSource)
	if err != nil {
		log.Println("DB error")
		panic(err)
	}
	engine.TZLocation, _ = time.LoadLocation("Asia/Shanghai")
	engine.ShowSQL(config.ShowSQL)
	engine.SetMaxIdleConns(config.MaxIdleConns)
	engine.SetMaxOpenConns(config.MaxOpenConns)
	DB = engine
	Sync()
}

func Sync() {

}
