package sqlite

import (
	"fmt"
	"sync"

	"github.com/bwcxyk/golibs/config"
	"github.com/bwcxyk/golibs/global/consts"
	"github.com/go-xorm/xorm"
	_ "github.com/mattn/go-sqlite3"
)

type GoLibSqlite = *xorm.EngineGroup

var sqliteOnce sync.Once
var sqliteNew GoLibSqlite

func NewSqlite() GoLibSqlite {
	sqliteOnce.Do(func() {
		var err error
		dsn := fmt.Sprintf("%s?loc=%s", config.GetSqlite().GetStorageFile(), config.GetSqlite().GetTimezone())
		sqliteNew, err = xorm.NewEngineGroup("sqlite3", []string{dsn})
		if config.Config.GetActive() == consts.EnvDev {
			sqliteNew.ShowSQL(true)
		}
		if err != nil {
			panic(err)
		}
	})
	return sqliteNew
}
