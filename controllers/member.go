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
  "fmt"
  "strconv"
  "strings"
  "regexp"
  _ "reflect"
  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/mysql"

  "../models"
  "../tools"
)

func member_init(c echo.Context, db *gorm.DB) models.Member {
  id       , _ := strconv.Atoi(c.QueryParam("id"))
  grade := models.Grade{Name: c.QueryParam("grade")}
  fmt.Println(grade.Name)
  db.First(&grade, "name = ?", grade.Name)
  // grade_id , _ := strconv.Atoi()
  return models.Member{id,c.QueryParam("name"),grade.Id,c.QueryParam("mail_address")}
}

func ErrorCheck(id int) bool{
  if id == 0 {
    return false
  }
  return true
}

func MemberAdd(db *gorm.DB) echo.HandlerFunc {
  return func(c echo.Context) error {
    member := member_init(c, db)
    grade := models.Grade{Name: c.QueryParam("grade")}
    db.First(&grade, "name = ?", grade.Name)
    if ErrorCheck(grade.Id) != true {
      em := tools.Error{Massege: "not exit grade"}
      return c.JSON(404, em)
    }
    db.Create(&member)
    name := c.QueryParam("project")
    if strings.Contains(name, "(") && strings.Contains(name, ")"){
        new_pro := strings.Split(name[1:len(name)-1], ",")
        for _, v := range new_pro{
          project := models.Project{}
          db.First(&project, "name = ?", v)
          if ErrorCheck(project.Id) != true {
            em := tools.Error{Massege: "not exit project"}
            return c.JSON(404, em)
          }
          db.Create(&models.Member_Project{Member_Id: member.Id, Project_Id: project.Id})
        }
      } else {
        project := models.Project{}
        db.First(&project, "name = ?", name)
        if ErrorCheck(grade.Id) != true {
          em := tools.Error{Massege: "not exit project"}
          return c.JSON(404, em)
        }
        db.Create(&models.Member_Project{Member_Id: member.Id, Project_Id: project.Id})
      }
      return c.JSON(http.StatusOK,tools.Error{Massege: "OK"})
  }
}

func MemberDelete(db *gorm.DB) echo.HandlerFunc {
  return func(c echo.Context) error {
    member := member_init(c, db)
    member_project := []models.Member_Project{}
    db.First(&member, "name = ?", member.Name)
    if ErrorCheck(member.Id) != true {
      em := tools.Error{Massege: "not exit Member"}
      return c.JSON(404, em)
    }
    db.Delete(&member)
    db.Find(&member_project, "member_id = ?", member.Id)
    for _, v := range member_project{
      if ErrorCheck(v.Member_Id) != true {
        em := tools.Error{Massege: "not exit Member"}
        return c.JSON(404, em)
      }
      db.Delete(&v)
    }
    return c.JSON(http.StatusOK,tools.Error{Massege: "OK"})
  }
}

func MemberUpdate(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		member := member_init(c, db)
		db.First(&member, "name = ?", member.Name)
		if ErrorCheck(member.Id) != true {
			em := tools.Error{Massege: "not exit member"}
			return c.JSON(404, em)
		}
		id, _ := strconv.Atoi(c.QueryParam("new_id"))
		grade := models.Grade{Name: c.QueryParam("new_grade")}
		db.First(&grade, "name = ?", grade.Name)
		new_data := models.Member{id, c.QueryParam("new_name"), grade.Id, c.QueryParam("new_mail_address")}
		new_project := models.Project{0, c.QueryParam("new_team")}
		old_data := db.Model(&member)
		if new_data.Name != "" {
			old_data.Update("name", new_data.Name)
		}
		if new_data.Grade_Id != 0 {
			old_data.Update("grade_id", new_data.Grade_Id)
		}
		if new_data.Mail_Address != "" {
			old_data.Update("mail_address", new_data.Mail_Address)
		}
		if new_project.Name != "" {
			name := new_project.Name
			member_project := []models.Member_Project{}
			project := models.Project{}
			if strings.Contains(name, "(") && strings.Contains(name, ")") {
				new_pro := strings.Split(name[1:len(name)-1], ",")
				for _, v := range new_pro {
					// project := Project{}
					db.First(&project, "name = ?", v)
					if ErrorCheck(project.Id) != true {
						em := tools.Error{Massege: "not exit project"}
						return c.JSON(404, em)
					}
					db.Create(&models.Member_Project{Member_Id: member.Id, Project_Id: project.Id})
				}
			} else {
				//project := Project{}
				db.First(&project, "name = ?", name)
				fmt.Println(project)
				if ErrorCheck(project.Id) != true {
					em := tools.Error{Massege: "not exit project"}
					return c.JSON(404, em)
				}
				db.Create(&models.Member_Project{Member_Id: member.Id, Project_Id: project.Id})
			}

			db.Find(&member_project, "member_id = ?", member.Id)
			for _, v := range member_project {
				if ErrorCheck(v.Member_Id) != true {
					em := tools.Error{Massege: "not exit Member"}
					return c.JSON(404, em)
				}
				if v.Project_Id != project.Id {
					fmt.Println(v)
					db.Delete(&v)
				}
			}
		}
		return c.JSON(http.StatusOK, tools.Error{Massege: "OK"})
	}
}

func ShowAddressCount(db *gorm.DB) echo.HandlerFunc {
  return func(c echo.Context) error {
    members := []models.Member{}
    count  := tools.Res_Json_Domain{}
    count.Domain_Type = map[string]int{}
    db.Find(&members)
    for _, v := range members {
      fmt.Println(v)
      r := regexp.MustCompile(`\s*@\s*`)
      domains := r.Split(v.Mail_Address, -1)
      fmt.Println(domains[1])
      domain := domains[1]
      if _, ok := count.Domain_Type[domain]; ok == true {
        count.Domain_Type[domain] += 1
      } else {
        count.Domain_Type[domain] =  1
      }
    }
    return c.JSON(http.StatusOK, count)
  }
}
