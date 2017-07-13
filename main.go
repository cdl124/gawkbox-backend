package main

import (
  "gawkbox-takehome/lib"

  "database/sql"
  _ "github.com/go-sql-driver/mysql"

  "fmt"
  "net/http"
  "github.com/spf13/viper"
)

var db *sql.DB
var err error

func main() {

  viper.SetConfigName("app")
  viper.AddConfigpath("/config/")

  err := viper.ReadInConfig()
  if err != nil {
    fmt.Println("Config file not found...")
  } else {
    mysqlconfig := viper.GetString("development.mysqlconfig")
    port := viper.GetString("development.port")

    fmt.Printf("\nDevelopment Config found:\n server = %s")
  }

  db, err = sql.Open("mysql", mysqlconfig)
  if err != nil {
    panic(err.Error())
  }
  defer db.Close()

  err = db.Ping()
  if err != nil {
    panic(err.Error())
  }

  http.HandleFunc("/signup", signupPage)
  http.HandleFunc("/login", loginPage)
  http.HandleFunc("/", homePage)
  http.ListenAndServe(port, nil)
}
