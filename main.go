package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-xorm/xorm"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

var Db *xorm.Engine

func main() {
	c := Init(os.Getenv("APP_ENV"))

	fmt.Println("Config===", c)
	var err error
	Db, err = InitDB(c.Database.Driver, c.Database.Connection)
	if err != nil {
		panic(err)
	}
	defer Db.Close()

	if err := InitTable(Db); err != nil {
		panic(err)
	}

	e := echo.New()

	e.GET("/ping", func(c echo.Context) error {
		return c.String(http.StatusOK, "pong")
	})

	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	e.Use(middleware.RequestID())
	e.Debug = c.Debug
	if err := e.Start(":" + c.HttpPort); err != nil {
		log.Println(err)
	}
}
