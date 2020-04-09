package main

import (
	"context"
	"os"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

var ctx context.Context

func TestMain(m *testing.M) {
	db := enterTest()
	code := m.Run()
	exitTest(db)
	os.Exit(code)
}

func enterTest() *xorm.Engine {
	c := Init(os.Getenv("APP_ENV"))
	var err error
	Db, err = xorm.NewEngine(c.Database.Driver, c.Database.Connection)
	if err != nil {
		panic(err)
	}
	// xormEngine.ShowSQL(true)
	// if err = DropTables(Db); err != nil {
	// 	panic(err)
	// }
	if err = InitTable(Db); err != nil {
		panic(err)
	}
	return Db
}

func exitTest(db *xorm.Engine) {
	//db.Close()
}
