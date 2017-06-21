package models

import (
  // フレームワーク関連パッケージ
  _ "github.com/labstack/echo"
  _ "github.com/labstack/echo/engine/standard"
  _ "github.com/labstack/echo/middleware"

  _ "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/mysql"
)

type Member struct {
  Id            int     `sql:"AUTO_INCREMENT"`
  Name          string  `json:name`
  Grade_Id      int     `json:"grade_id"`
  Mail_Address  string  `json:"mail_address"`
}
