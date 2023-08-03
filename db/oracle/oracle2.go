/*
@Author : YaoKun
@Time : 2023/8/3 10:02
*/

package oracle

import (
	"database/sql"
	"github.com/bwcxyk/golibs/config"
	go_ora "github.com/sijms/go-ora/v2"
)

func connectToDatabase() (*sql.DB, error) {
	dsn := go_ora.BuildUrl(
		config.Config.GetOracle().GetHost(),
		config.Config.GetOracle().GetPort(),
		config.Config.GetOracle().GetDbname(),
		config.Config.GetOracle().GetUsername(),
		config.Config.GetOracle().GetPassword(),
		nil)

	// 使用连接字符串打开数据库连接
	db, err := sql.Open("oracle", dsn)
	if err != nil {
		return nil, err
	}

	// 设置连接池的最大连接数和空闲连接数
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)

	return db, nil
}

func queryDatabase(db *sql.DB, query string) {
	// 执行查询
	rows, _ := db.Query(query)
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
		}
	}(rows)
}
