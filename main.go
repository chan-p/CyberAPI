package main

import (
  // フレームワーク関連パッケージ
  "github.com/labstack/echo"
  "github.com/labstack/echo/engine/standard"
  "github.com/labstack/echo/middleware"
  
  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/mysql"

  "./controllers"
)

func db_connect() *gorm.DB{
  db,err := gorm.Open("mysql","root:tomohi6@tcp(db01.wsl.mind.meiji.ac.jp:3306)/wsl_member_table_A")
  if err != nil {
    panic(err.Error())
  }
  return db
}
func main(){

  e := echo.New()
  db := db_connect()
  db.DB()
  defer db.Close()
  // ミドルウェアの使用機能
  e.Use(middleware.Logger())
  e.Use(middleware.Recover())

  // ルーティング
  e.Get("/api/data/member_add"       , controllers.MemberAdd(db)       )
  e.Get("/api/data/member_delete"    , controllers.MemberDelete(db)    )
  e.Get("/api/data/member_update"    , controllers.MemberUpdate(db)    )
  e.Get("/api/show/grade_all"        , controllers.ShowGradAlle(db)    )
  e.Get("/api/show/team_member_count", controllers.ShowTeamCount(db)   )
  e.Get("/api/show/address_count"    , controllers.ShowAddressCount(db))

  // サーバー構築 ポート8000
  e.Run(standard.New(":50000"))
}
