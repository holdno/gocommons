package mysql

import (
	"database/sql"

	"github.com/go-sql-driver/mysql"
)

type MysqlConfig struct {
	User     string
	Password string
	Addr     string
	DBName   string
}

var mysqlInstance *sql.DB

func Init(conf *MysqlConfig) {
	var (
		err error
		mc  mysql.Config
	)
	mc = mysql.Config{
		User:                 conf.User,
		Passwd:               conf.Password,
		Net:                  "tcp",
		Addr:                 conf.Addr,
		DBName:               conf.DBName,
		AllowNativePasswords: true,
	}
	if mysqlInstance, err = sql.Open("mysql", mc.FormatDSN()); err != nil {
		panic("Mysql init error" + err.Error())
	}
}

func GetDB() *sql.DB {
	return mysqlInstance
}
