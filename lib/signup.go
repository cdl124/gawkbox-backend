package lib

import "database/sql"
import _ "github.com/go-sql-driver/mysql"
import "golang.org/x/crypto/bcrypt"

import "net/http"

//var db *sql.DB
//var err Error

func SignupPage(res http.ResponseWriter, req *http.Request) {
  if req.Method != "POST" {
    http.ServeFile(res, req, "signup.html")
    return
  }

  username := req.FormValue("username")
  password := req.FormValue("password")

  var user string

  err := db.QueryRow("SELECT username FROM users WHERE username=?", username).Scan(&user)

  switch {
  case err == sql.ErrNoRows: // the username is available
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    if err != nil {
      http.Error(res, "Server error, unable to create your account.", 500)
      return
    }
    _, err = db.Exec("INSERT INTO users(username, password) VALUES(?, ?)", username, hashedPassword)
    if err != nil {
      http.Error(res, "Server error, unable to create your account.", 500)
      return
    }

    res.Write([]byte("Username created!"))
    http.ServeFile(res, req, "login.html")

  case err != nil:
    http.Error(res, "Server error, unable to create your account.", 500)
    return

  default:
    http.Redirect(res, req, "/", 301)
  }
}
