package tool

import (
  "net/http"
  "github.com/labstack/echo"
  "database/sql"
  "fmt"
  _ "github.com/go-sql-driver/mysql"
)

func get_connect() *sql.DB{
  db,err := sql.Open("mysql","root:tomohi6@tcp(192.168.60.55:3306)/social-app")
  if err != nil {
    panic(err.Error())
  }
  return db
}

func Res_mysql() echo.HandlerFunc {
  return func(c echo.Context) error {
    db := get_connect()
    rows, err := db.Query("select * from test")
    fmt.Println(rows)
    if err != nil {
      return err
    }
    colum,err := rows.Columns()
    values := make([]sql.RawBytes,len(colum))
    scanArgs := make([]interface{},len(values))
    for i := range values {
      scanArgs[i] = &values[i]
    }
    var value string
    for rows.Next(){
      err= rows.Scan(scanArgs...)
      if err != nil {
        return err
      }
      for _,col := range values {
        if col == nil {
          value = "NULL"
        } else {
          value = string(col)
        }
      }
    }
    return c.String(http.StatusOK,value)
  }
}
