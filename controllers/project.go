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

func ShowTeamCount(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		project := models.Project{}
		db.First(&project, "name = ?", c.QueryParam("team"))
		if ErrorCheck(project.Id) != true {
			em := tools.Error{Massege: "not exit team"}
			return c.JSON(404, em)
		}
		member_project := []models.Member_Project{}
		db.Find(&member_project, "project_id = ?", project.Id)
		response := tools.Res_Json{
			ProjectMemberCount: len(member_project),
		}
		return c.JSON(http.StatusOK, response)
	}
}
