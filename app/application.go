package main

import (
	// フレームワーク関連パッケージ
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	"github.com/labstack/echo/middleware"

	// データベース関連パッケージ
	_ "database/sql"

	_ "github.com/go-sql-driver/mysql"

	"fmt"
	"net/http"
	_ "reflect"
	"regexp"
	"strconv"
	"strings"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Member struct {
	Id           int    `sql:"AUTO_INCREMENT"`
	Name         string `json:"name"`
	Grade_Id     int    `json:"grade_id"`
	Mail_Address string `json:"mail_address"`
}

type Grade struct {
	Id   int    `sql:"AUTO_INCREMENT"`
	Name string `json:"name"`
}

type Project struct {
	Id   int    `sql:"AUTO_INCREMENT"`
	Name string `json:"name"`
}

type Member_Project struct {
	Id         int `sql:"AUTO_INCREMENT"`
	Member_Id  int `json:"member_id"`
	Project_Id int `json:"project_id"`
}

type Res_Json struct {
	ProjectMemberCount int `json:"projectmembercount"`
}

type Res_Json_Domain struct {
	Domain_Type map[string]int `json:"domain_type"`
}

type Res_Json_Member_List struct {
	Name []string `json:"name"`
}

type Error struct {
	Massege string `json:"massege"`
}

func db_connect() *gorm.DB {
	db, err := gorm.Open("mysql", "root:tomohi6@tcp(db01.wsl.mind.meiji.ac.jp:3306)/wsl_member_table_A")
	if err != nil {
		panic(err.Error())
	}
	return db
}

func member_init(c echo.Context, db *gorm.DB) Member {
	id, _ := strconv.Atoi(c.QueryParam("id"))
	grade := Grade{Name: c.QueryParam("grade")}
	fmt.Println(grade.Name)
	db.First(&grade, "name = ?", grade.Name)
	// grade_id , _ := strconv.Atoi()
	return Member{id, c.QueryParam("name"), grade.Id, c.QueryParam("mail_address")}
}

func ErrorCheck(id int) bool {
	if id == 0 {
		return false
	}
	return true
}

func MemberAdd(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		member := member_init(c, db)
		grade := Grade{Name: c.QueryParam("grade")}
		fmt.Println(grade)
		fmt.Println(c.QueryParam("grade"))
		db.First(&grade, "name = ?", grade.Name)
		if ErrorCheck(grade.Id) != true {
			em := Error{Massege: "not exit grade"}
			return c.JSON(404, em)
		}
		db.Create(&member)
		name := c.QueryParam("project")
		if strings.Contains(name, "(") && strings.Contains(name, ")") {
			new_pro := strings.Split(name[1:len(name)-1], ",")
			for _, v := range new_pro {
				project := Project{}
				db.First(&project, "name = ?", v)
				if ErrorCheck(project.Id) != true {
					em := Error{Massege: "not exit project"}
					return c.JSON(404, em)
				}
				db.Create(&Member_Project{Member_Id: member.Id, Project_Id: project.Id})
			}
		} else {
			project := Project{}
			db.First(&project, "name = ?", name)
			if ErrorCheck(grade.Id) != true {
				em := Error{Massege: "not exit project"}
				return c.JSON(404, em)
			}
			db.Create(&Member_Project{Member_Id: member.Id, Project_Id: project.Id})
		}
		return c.JSON(http.StatusOK, Error{Massege: "OK"})
	}
}

func MemberDelete(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		member := member_init(c, db)
		member_project := []Member_Project{}
		db.First(&member, "name = ?", member.Name)
		if ErrorCheck(member.Id) != true {
			em := Error{Massege: "not exit Member"}
			return c.JSON(404, em)
		}
		db.Delete(&member)
		db.Find(&member_project, "member_id = ?", member.Id)
		for _, v := range member_project {
			if ErrorCheck(v.Member_Id) != true {
				em := Error{Massege: "not exit Member"}
				return c.JSON(404, em)
			}
			db.Delete(&v)
		}
		return c.JSON(http.StatusOK, Error{Massege: "OK"})
	}
}

func MemberUpdate(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		member := member_init(c, db)
		db.First(&member, "name = ?", member.Name)
		if ErrorCheck(member.Id) != true {
			em := Error{Massege: "not exit member"}
			return c.JSON(404, em)
		}
		id, _ := strconv.Atoi(c.QueryParam("new_id"))
		grade := Grade{Name: c.QueryParam("new_grade")}
		db.First(&grade, "name = ?", grade.Name)
		new_data := Member{id, c.QueryParam("new_name"), grade.Id, c.QueryParam("new_mail_address")}
		new_project := Project{0, c.QueryParam("new_project")}
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
			member_project := []Member_Project{}
			if strings.Contains(name, "(") && strings.Contains(name, ")") {
				new_pro := strings.Split(name[1:len(name)-1], ",")
				for _, v := range new_pro {
					project := Project{}
					db.First(&project, "name = ?", v)
					if ErrorCheck(project.Id) != true {
						em := Error{Massege: "not exit grade"}
						return c.JSON(404, em)
					}
					db.Create(&Member_Project{Member_Id: member.Id, Project_Id: project.Id})
				}
			} else {
				project := Project{}
				db.First(&project, "name = ?", name)
				fmt.Println(project)
				if ErrorCheck(project.Id) != true {
					em := Error{Massege: "not exit grade"}
					return c.JSON(404, em)
				}
				db.Create(&Member_Project{Member_Id: member.Id, Project_Id: project.Id})
			}
			db.Find(&member_project, "member_id = ?", member.Id)
			for _, v := range member_project {
				if ErrorCheck(v.Member_Id) != true {
					em := Error{Massege: "not exit Member"}
					return c.JSON(404, em)
				}
				db.Delete(&v)
			}
		}
		return c.JSON(http.StatusOK, Error{Massege: "OK"})
	}
}

func ShowGradAlle(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		grade := Grade{}
		db.First(&grade, "name = ?", c.QueryParam("grade"))
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

func ShowTeamCount(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		project := Project{}
		db.First(&project, "name = ?", c.QueryParam("team"))
		if ErrorCheck(project.Id) != true {
			em := Error{Massege: "not exit team"}
			return c.JSON(404, em)
		}
		member_project := []Member_Project{}
		db.Find(&member_project, "project_id = ?", project.Id)
		response := Res_Json{
			ProjectMemberCount: len(member_project),
		}
		return c.JSON(http.StatusOK, response)
	}
}

func ShowAddressCount(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		members := []Member{}
		count := Res_Json_Domain{}
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
				count.Domain_Type[domain] = 1
			}
		}
		return c.JSON(http.StatusOK, count)
	}
}

func main() {

	e := echo.New()
	db := db_connect()
	db.DB()
	defer db.Close()
	// ミドルウェアの使用機能
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// モデル
	e.Get("/api/data/member_add", MemberAdd(db))
	e.Get("/api/data/member_delete", MemberDelete(db))
	e.Get("/api/data/member_update", MemberUpdate(db))
	e.Get("/api/show/grade_all", ShowGradAlle(db))
	e.Get("/api/show/team_member_count", ShowTeamCount(db))
	e.Get("/api/show/address_count", ShowAddressCount(db))

	// サーバー構築 ポート8000
	e.Run(standard.New(":1323"))
}
