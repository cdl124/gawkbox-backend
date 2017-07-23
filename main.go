package main

import (
  "gawkbox-takehome/lib"
  "github.com/spf13/viper"
  "database/sql"
  _ "github.com/go-sql-driver/mysql"

  "fmt"
  "net/http"
)

var db *sql.DB
var err error

func main() {

  viper.SetConfigName("app")
  viper.AddConfigPath("config")
  viper.AddConfigPath(".")

  err := viper.ReadInConfig()
  if err != nil {
    fmt.Println("Config file not found...")
  } else {
    fmt.Printf("\nDevelopment Config found:\n server = ", viper.GetString("development.port"))
  }

  db, err = sql.Open("mysql", viper.GetString("development.mysqlconfig"))
  if err != nil {
    panic(err.Error())
  }
  defer db.Close()

  err = db.Ping()
  if err != nil {
    panic(err.Error())
  }

  http.HandleFunc("/signup", lib.SignupPage)
  http.HandleFunc("/login", lib.LoginPage)
  http.HandleFunc("/", lib.HomePage)
  http.ListenAndServe(viper.GetString("development.port"), nil)
}
