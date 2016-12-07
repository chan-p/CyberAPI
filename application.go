package main

import (
  // フレームワーク関連パッケージ
  "github.com/labstack/echo"
  "github.com/labstack/echo/engine/standard"
  "github.com/labstack/echo/middleware"

  // データベース関連パッケージ
  "database/sql"
  _ "github.com/go-sql-driver/mysql"

  // ディレクトリ
  "./tool"
  _ "./model"
  _ "fmt"
)

func db_connect() *sql.DB {
  db,err := sql.Open("mysql","root:tomonori@tcp(52.196.55.156:3306)/social_app")

  if err != nil {
    panic(err.Error())
  }
  return db
}

func main(){

  e := echo.New()
  db := db_connect()
  defer db.Close()
  // ミドルウェアの使用機能
  e.Use(middleware.Logger())
  e.Use(middleware.Recover())

  // モデル
  e.Get("/json",tool.Res_json(db))

  //サーバー構築 ポート8000
  e.Run(standard.New(":8000"))
}
