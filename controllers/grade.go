package controllers

import (
  // フレームワーク関連パッケージ
  "github.com/labstack/echo"
  _ "github.com/labstack/echo/engine/standard"
  _ "github.com/labstack/echo/middleware"

  // データベース関連パッケージ
  _ "database/sql"
  _ "github.com/go-sql-driver/mysql"

  "net/http"
  _ "fmt"
  _ "strconv"
  _ "strings"
  _ "regexp"
  _ "reflect"
  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/mysql"

  "../models"
  "../tools"
)

func ShowGradAlle(db *gorm.DB) echo.HandlerFunc {
  return func(c echo.Context) error {
    grade  := models.Grade{}
    db.First(&grade , "name = ?", c.QueryParam("grade"))
    if ErrorCheck(grade.Id) != true {
      em := tools.Error{Massege: "not exit grade"}
      return c.JSON(404, em)
    }
    members := []models.Member{}
    db.Find(&members, "grade_id = ?", grade.Id)
    list := tools.Res_Json_Member_List{}
    for _, v := range members {
      list.Name = append(list.Name, v.Name)
    }
    return c.JSON(http.StatusOK, list)
  }
}
