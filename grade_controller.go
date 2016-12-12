package main

import (
  // フレームワーク関連パッケージ
  "github.com/labstack/echo"
  "github.com/labstack/echo/engine/standard"
  "github.com/labstack/echo/middleware"

  // データベース関連パッケージ
  _ "database/sql"
  _ "github.com/go-sql-driver/mysql"

  "net/http"
  "fmt"
  "strconv"
  "strings"
  "regexp"
  _ "reflect"
  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/mysql"

  "./model"
)

func ShowGradAlle(db *gorm.DB) echo.HandlerFunc {
  return func(c echo.Context) error {
    grade  := Grade{}
    db.First(&grade , "name = ?", c.QueryParam("grade"))
    if ErrorCheck(grade.Id) != true {
      em := Error{Massege: "not exit grade"}
      return c.JSON(404, em)
    }
    members := []Member{}
    db.Find(&members, "grade_id = ?", grade.Id)
    list := Res_Json_Member_List{}
    for _, v := range members {
      list.Name = append(list.Name, v.Name)
    }
    return c.JSON(http.StatusOK, list)
  }
}
