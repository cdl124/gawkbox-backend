package lib

import "net/http"

func HomePage(res http.ResponseWriter, req *http.Request) {
  http.ServeFile(res, req, "index.html")
}
