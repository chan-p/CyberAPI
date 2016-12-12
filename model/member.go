package model

import (
  // フレームワーク関連パッケージ
  "github.com/labstack/echo"
  "github.com/labstack/echo/engine/standard"
  "github.com/labstack/echo/middleware"

  "net/http"
  "fmt"
  "strconv"
  "strings"
  "regexp"
  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/mysql"
)

type Member struct {
  Id            int     `sql:"AUTO_INCREMENT"`
  Name          string  `json:name`
  Grade_Id      int     `json:"grade_id"`
  Mail_Address  string  `json:"mail_address"`
}
