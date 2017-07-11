package login

import "database/sql"
import _ "github.com/go-sql-driver/mysql"

import "golang.org/x/crypto/bcrypt"

import "net/http"

var db *sql.DB
var err Error

func loginPage(res http.ResponseWriter, req *http.Request) {
  if req.Method != "POST" {
    http.ServeFile(res, req, "login.html")
    return
  }

  username := req.FormValue("username")
  password := req.FormValue("password")

  var databaseUsername string
  var databasePassword string

  err := db.QueryRow("SELECT username, password FROM users WHERE username=?", username).Scan(&databaseUsername, &databasePassword)

  if err != nil {
    http.Redirect(res, req, "/", 301)
    return
  }

  http.ServeFile(res, req, "index.html")
}
