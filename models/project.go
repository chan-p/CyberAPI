package models

import (
  _ "github.com/jinzhu/gorm/dialects/mysql"
)

type Project struct {
  Id            int     `sql:"AUTO_INCREMENT"`
  Name          string  `json:name`
}


