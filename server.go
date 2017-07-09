package main

import "./lib/signup"
import "./lib/login"

import "database/sql"
import _ "github.com/go-sql-driver/mysql"

import "net/http"

var db *sql.DB
var err error

func main() {
  db, err = sql.Open("mysql", "root:govithecactus@/highfive")
  if err != nil {
    panic(err.Error())
  }
  defer db.Close()

  err = db.Ping()
  if err != nil {
    panic(err.Error())
  }

  http.ListenAndServe(":8080", nil)
}
