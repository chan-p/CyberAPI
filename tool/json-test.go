package tool

import (
  "net/http"
  "github.com/labstack/echo"
  "time"
)

type Data struct {
  Id  int       `json:"id"`
  Title string    `json:"title"`
  CreatedAt time.Time `json:"created_at"`
  Query string `json:"query"`
}

func Res_json() echo.HandlerFunc {
  return func(c echo.Context) error {
    loc, _ := time.LoadLocation("Asia/Tokyo")
    d := &Data{
      Id: 1,
      Title: "自動変更完成pull もう一度確認 再度確認",
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
