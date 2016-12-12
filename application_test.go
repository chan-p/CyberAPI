package test

import (
  _ "bytes"

  "io/ioutil"
  "net/http"
  "testing"
)

func TestRequest(t *testing.T) {

  url := "http://localhost:1323/api"

  //postrequest作成
  req, err := http.NewRequest("GET", url + "/show/address_count", nil)
  //httpクライアント
  client := &http.Client{}
  //実行
  resp, err := client.Do(req)
  if err != nil {
    panic(err)
  }

  //ステータスコード確認
  if resp.StatusCode != 200 {
    t.Error(resp.StatusCode)
    return
  }

  //レスポンスBODY取得
  body, _ := ioutil.ReadAll(resp.Body)

  actual := string(body)
  expected := `{"Domain_Type":{"gmail.com":21,"icloud.com":1,"outlook.jp":1,"yahoo.co.jp":2}}`

  if actual != expected {
    t.Error("response error")
  }
  return
}
