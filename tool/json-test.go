package tool

import (
	//apiパッケージ
	"github.com/labstack/echo"
	"net/http"

	//mysqlパッケージ
	"database/sql"
	_ "github.com/go-sql-driver/mysql"

	//標準パッケージ
	_ "fmt"
  "time"
	"strconv"
  _ "strings"
)

type Data struct {
  Id        int       `json:"id"`
  Title     string    `json:"title"`
  CreatedAt time.Time `json:"created_at"`
  Query     string    `json:"query"`
}

func db_test(db *sql.DB) []string{
  query := "select *  from User limit 1"
  return extract_from_db(db, query)
}

func Res_json(db *sql.DB) echo.HandlerFunc {
  return func(c echo.Context) error {
    loc, _ := time.LoadLocation("Asia/Tokyo")
    data := db_test(db)
    id, _ := strconv.Atoi(data[0])
    d := &Data{
      Id: id,
      Title: data[1],
      CreatedAt: time.Date(2014, 8, 25, 0, 0, 0, 0, loc),
      Query: c.FormValue("name"),
    }
    //bytes, _ := json.Marshal(d)
    if err := c.Bind(d); err != nil{
      return err
    }
    return c.JSON(http.StatusOK,d)
  }
}
