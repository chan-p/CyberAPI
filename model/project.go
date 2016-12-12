package model

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
)

type Project struct {
  Id            int     `sql:"AUTO_INCREMENT"`
  Name          string  `json:name`
}


