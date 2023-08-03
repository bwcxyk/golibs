/*
@Author : YaoKun
@Time : 2022/9/30 13:25
*/

package oracle

import (
	"database/sql"
	"log"
	"testing"
)

func TestDb2(t *testing.T) {
	// 连接数据库
	db, err := ConnectToDatabase()
	if err != nil {
		log.Fatal(err)
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
		}
	}(db)

	QueryDatabase(db, "select sysdate from dual")
}
