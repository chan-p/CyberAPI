package model

import (
	//apiパッケージ
	"github.com/labstack/echo"
	"net/http"
  _ "github.com/labstack/echo/engine/standard"
  _ "github.com/labstack/echo/middleware"


	//mysqlパッケージ
  "github.com/jinzhu/gorm"
	_ "database/sql"
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
  list := []string{"name","casts","crews"}
  for i := 0; i < 4;i = i + 1 {
    _ ,OK := query[list(i)]
    if OK != nil{
      queray += list[i]
    }
  }
  fmt.Println(queray)
}

func Echo_add(db *sql.DB) echo.HandlerFunc {
  return func(c echo.Context) error {
    data := _init(c)
    fmt.Println(data)
    status := data.get_data()
    return c.JSON(http.StatusOK,status)
  }
}
