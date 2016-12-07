package model

import (
	//apiパッケージ
	"github.com/labstack/echo"
	"net/http"
  _ "github.com/labstack/echo/engine/standard"
  _ "github.com/labstack/echo/middleware"


	//mysqlパッケージ
	"database/sql"
	_ "github.com/go-sql-driver/mysql"

	//標準パッケージ
	"fmt"
	_ "strconv"
  _ "strings"
)
type res_json struct {
	Data   []res_data
}

type res_data struct {
  Id    int     `json:id`
	Name  string  `json:name`
	Casts string  `json:casts`
	Crews string  `json:crews`
}

type query_data struct {
	name  string
	casts string
	crews string
  op    string
}

func series_init(c echo.Context) query_data {
	return query_data{c.QueryParam("name"), c.QueryParam("casts"), c.QueryParam("crews"), c.QueryParam("op")}
}

func (query query_data) get_data() string {
  queray := "select id, content, casts ,crews from series"
  return queray
  // for i := 0; i < 
}

func Echo_api_no1(db *sql.DB) echo.HandlerFunc {
  return func(c echo.Context) error {
    data := series_init(c)
    fmt.Println(data)
    status := data.get_data()
    return c.JSON(http.StatusOK,status)
  }
}