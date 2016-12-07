package tool

import (
  "net/http"
  "github.com/labstack/echo"
  _ "database/sql"
  "fmt"
  _ "github.com/go-sql-driver/mysql"
)

type user_data struct {
  email string
  year string
  month string
}

func initation(c echo.Context) user_data{
  return user_data{c.QueryParam("email"),c.QueryParam("year"),c.QueryParam("month")}
}

func (user user_data) get_event() string{
  return user.email +":"+string(user.year)+":"+string(user.month)
}

func Echo_event() echo.HandlerFunc {
  return func(c echo.Context) error {
    user := initation(c)
    fmt.Println(user)
    a := user.get_event()
    return c.String(http.StatusOK,a)
  }
}
