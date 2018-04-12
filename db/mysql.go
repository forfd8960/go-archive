package db

import (
	"database/sql"

	"github.com/forfd8960/go-archive/conf"
)

var MySql *sql.DB

func InitMysql(c *conf.MySQLConfig) error {
	pool, err := sql.Open("mysql", c.Dsn)
	if err != nil {
		return err
	}

	if c.MaxIdle > 0 {
		pool.SetMaxIdleConns(c.MaxIdle)
	}
	if c.MaxConn > 0 {
		pool.SetMaxOpenConns(c.MaxConn)
	}

	MySql = pool
	return nil
}
