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
	)
	if mysqlInstance, err = sql.Open("mysql", mysql.Config{
		User:                 conf.User,
		Passwd:               conf.Password,
		Net:                  "tcp",
		Addr:                 conf.Addr,
		DBName:               conf.DBName,
		AllowNativePasswords: true,
	}.FormatDSN()); err != nil {
		panic("Mysql init error" + err.Error())
	}
}

func GetDB() *sql.DB {
	return mysqlInstance
}
