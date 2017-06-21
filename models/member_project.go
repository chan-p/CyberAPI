package models

import (
  _ "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/mysql"
)

type Member_Project struct {
  Id            int     `sql:"AUTO_INCREMENT"`
  Member_Id     int     `json:member_id`
  Project_Id    int     `json:project_id`
}
